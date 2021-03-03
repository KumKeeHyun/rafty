package raft

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	pb "github.com/KumKeeHyun/rafty/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRaftyNode(cfg *Config) (*RaftyNode, error) {
	clusterConfig, err := newClusterConfig(cfg)
	if err != nil {
		return nil, err
	}

	nodeWG := &sync.WaitGroup{}
	nodeCtx, nodeCancel := context.WithCancel(context.Background())

	nodeContext := newNodeCtx()
	logStorage := newLogStorage()

	nodeContext.startNodeCtx(nodeCtx, nodeWG)
	logStorage.startLogStorage(nodeCtx, nodeWG)

	return &RaftyNode{
		config:                   clusterConfig,
		nodeContext:              nodeContext,
		logStorage:               logStorage,
		nodeWG:                   nodeWG,
		nodeCancel:               nodeCancel,
		ctxCancel:                func() {},
		UnimplementedRaftyServer: pb.UnimplementedRaftyServer{},
	}, nil
}

type RaftyNode struct {
	config      *clusterConfig
	nodeContext *nodeCtx    // term, votedFor, state
	logStorage  *logStorage // log entries, snapshot

	nodeWG     *sync.WaitGroup    // RaftyNode 하위에 실행된 고루틴 담당
	nodeCancel context.CancelFunc // RaftyNode 하위에 실행된 고루틴에 종료 신호를 전달하는 함수
	ctxCancel  context.CancelFunc // Candidate, Leader 상태에서 실행한 고루틴을 중지시키는 함수. becomeFollower 에서 호출

	pb.UnimplementedRaftyServer
}

// startElection 선거 타이머가 끝나면 클러스터에 리더가 없다고 판단하고 투표 시작
// Term 증가, 자신에게 투표한 뒤 다른 노드들에게 RequestVote RPC 호출
func (n *RaftyNode) startElection() {
	var (
		currentTerm  uint64
		lastLogIndex uint64
		lastLogTerm  uint64
		votes        int64 = 1
	)

	n.ctxCancel()
	ctx, ctxCancel := context.WithTimeout(context.Background(), electionTimeoutDuration)
	n.ctxCancel = ctxCancel

	n.nodeContext.Do(func(nc *nodeCtx) {
		nc.state = stateCandidate
		nc.term++
		currentTerm = nc.term
		nc.votedFor = n.config.id
		nc.currentLeader = -1
	})
	n.logStorage.Do(func(ls *logStorage) {
		lastLogIndex, lastLogTerm = ls.lastLogIndexAndTerm()
	})

	for _, peerNode := range n.config.peers {
		go func(p *peer) {
			resp, err := p.cli.RequestVote(ctx, &pb.RequestVoteReq{
				Term:         currentTerm,
				CandidateID:  n.config.id,
				LastLogIndex: lastLogIndex,
				LastLogTerm:  lastLogTerm,
			}, grpc.WaitForReady(true))
			if err != nil {
				return
			}

			isCandidate := true
			n.nodeContext.Do(func(nc *nodeCtx) { isCandidate = (nc.state == stateCandidate) })
			if !isCandidate {
				return
			}

			respTerm := resp.GetTerm()
			if currentTerm < respTerm {
				n.becomeFollower(respTerm)
				return
			} else if currentTerm == respTerm && resp.GetVoteGranted() {
				v := atomic.AddInt64(&votes, 1)
				if n.config.isQuorum(v) {
					n.startLeader()
					return
				}
			}
		}(peerNode)
	}
}

func (n *RaftyNode) startLeader() {
	getTimer().stop()
	n.ctxCancel()
	leaderCtx, leaderCancel := context.WithCancel(context.Background())
	n.ctxCancel = leaderCancel
	n.nodeContext.Do(func(nc *nodeCtx) {
		nc.state = stateLeader
		nc.currentLeader = n.config.id
	})

	var lastLogIndex uint64
	n.logStorage.Do(func(ls *logStorage) { lastLogIndex, _ = ls.lastLogIndexAndTerm() })
	logReplication := newLogReplication(n.config.peers, lastLogIndex) // 로그 복제를 위한 nextIndex, matchIndex 초기화

	go func() {
		var (
			heartbeatTicker = time.NewTicker(sendHeartbeatsDuration)
			hbCtx           context.Context
			hbCancel        context.CancelFunc = func() {}
		)
		defer heartbeatTicker.Stop()

		for {
			select {
			case <-heartbeatTicker.C:
				isLeader := true
				n.nodeContext.Do(func(nc *nodeCtx) { isLeader = (nc.state == stateLeader) })
				if !isLeader {
					hbCancel()
					return
				}
				// TODO: logReplication 검사해서 commitIndex 업데이트, stateMachine 업데이트, 필요하다면 snapshot 업데이트
				hbCtx, hbCancel = context.WithTimeout(context.Background(), sendHeartbeatsDuration)
				n.sendHeartbeats(hbCtx, logReplication)

			case <-leaderCtx.Done():
				hbCancel()
				return
			}
		}
	}()
}

func (n *RaftyNode) sendHeartbeats(ctx context.Context, logReplica *logReplication) {
	var (
		currentTerm uint64 = 0
	)
	n.nodeContext.Do(func(nc *nodeCtx) { currentTerm = nc.term })

	for _, peerNode := range n.config.peers {
		go func(p *peer) {
			var (
				nextIdx     uint64      = 0
				commitIdx   uint64      = 0
				prevLogIdx  uint64      = 0
				prevLogTerm uint64      = 0
				entires     []*pb.Entry = nil
			)

			nextIdx = logReplica.nextIndex[p.id] // 논리적인 index
			n.logStorage.Do(func(ls *logStorage) {
				ni := ls.getEntryIndex(nextIdx) // 메모리에서 관리하고 있는 entries의 index로 변환
				if ni >= 0 {                    // ni가 음수라면 snapshot에 포함된 entry임.
					commitIdx = ls.commitIndex
					entires = ls.entries[ni:]
					if ni > 0 {
						prevLogIdx = ls.entries[ni-1].Index
						prevLogTerm = ls.entries[ni-1].Term
					} else {
						prevLogIdx = ls.snapshotMetadata.lastLogIndex
						prevLogTerm = ls.snapshotMetadata.lastLogTerm
					}
				}
			})
			if entires == nil { // nextIdx가 snapshot에 속함. installSnapshot RPC 호출
				// TODO: snapshot 보내기. 스냅샷 보내기를 전담하는 객체, 고루틴이 있어야 할 것 같음
				return
			}

			resp, err := p.cli.AppendEntries(ctx, &pb.AppendEntriesReq{
				Term:         currentTerm,
				Leader:       n.config.id,
				PrevLogIndex: prevLogIdx,
				PrevLogTerm:  prevLogTerm,
				LeaderCommit: commitIdx,
				Entries:      entires,
			}, grpc.WaitForReady(true))
			if err != nil {
				return
			}

			respTerm := resp.GetTerm()
			if respTerm > currentTerm {
				n.becomeFollower(respTerm)
				return
			}

			isLeader := true
			n.nodeContext.Do(func(nc *nodeCtx) { isLeader = (nc.state == stateLeader) })
			if isLeader && respTerm == currentTerm {
				if resp.Success { // 성공적으로 복제되었다면 nextIndex, matchIndex 업데이트
					logReplica.nextIndex[p.id] = nextIdx + uint64(len(entires))
					logReplica.matchIndex[p.id] = logReplica.nextIndex[p.id] - 1
				} else { // follower의 log entries가 nextIdx보다 뒤쳐저 있다면 nextIndex 조정
					logReplica.nextIndex[p.id] = nextIdx - resp.GetLag()
				}
			}
		}(peerNode)
	}
}

func (n *RaftyNode) becomeFollower(term uint64) {
	defer getTimer().reset()
	getTimer().stop()
	n.ctxCancel()
	n.nodeContext.Do(func(nc *nodeCtx) {
		nc.state = stateFollower
		nc.term = term
		nc.votedFor = 0
		nc.currentLeader = -1
	})
}

// ============= grpc ==============

func (n *RaftyNode) RequestVote(ctx context.Context, in *pb.RequestVoteReq) (*pb.RequestVoteResp, error) {
	var (
		resp = &pb.RequestVoteResp{
			Term:        0,
			VoteGranted: false,
		}
		lastLogIndex, lastLogTerm uint64
	)

	n.logStorage.Do(func(ls *logStorage) { lastLogIndex, lastLogTerm = ls.lastLogIndexAndTerm() })
	n.nodeContext.Do(func(nc *nodeCtx) { resp.Term = nc.term })
	reqTerm := in.GetTerm()
	if reqTerm > resp.Term {
		n.becomeFollower(reqTerm)
		resp.Term = reqTerm
	}

	reqIndex := in.GetLastLogIndex()
	reqID := in.GetCandidateID()
	n.nodeContext.Do(func(nc *nodeCtx) {
		if reqTerm == nc.term &&
			(nc.votedFor == 0 || nc.votedFor == reqID) &&
			(reqTerm > lastLogTerm || (reqTerm == lastLogTerm && reqIndex >= lastLogIndex)) {
			resp.VoteGranted = true
			nc.votedFor = reqID
		}
	})
	if resp.VoteGranted {
		getTimer().reset()
	}

	return resp, nil
}

func (n *RaftyNode) AppendEntries(ctx context.Context, in *pb.AppendEntriesReq) (*pb.AppendEntriesResp, error) {
	var (
		resp = &pb.AppendEntriesResp{
			Term:    0,
			Success: false,
			Lag:     0,
		}
		lastLogIndex, lastLogTerm uint64
		prevLogTerm               uint64 = 0
	)

	n.nodeContext.Do(func(nc *nodeCtx) { resp.Term = nc.term })
	reqTerm := in.GetTerm()
	if reqTerm > resp.Term {
		n.becomeFollower(reqTerm)
		resp.Term = reqTerm
	}

	if resp.Term == reqTerm {
		isFollower := true
		n.nodeContext.Do(func(nc *nodeCtx) { isFollower = (nc.state == stateFollower) })
		if !isFollower {
			n.becomeFollower(reqTerm)
		}
		getTimer().reset()

		n.logStorage.Do(func(ls *logStorage) {
			lastLogIndex, lastLogTerm = ls.lastLogIndexAndTerm()
			prevIndex := ls.getEntryIndex(in.GetPrevLogIndex()) // 이부분부터 in.entries가 follower의 commit된 entries를 가리킬 수 있는지 모르겠음.
			if prevIndex >= 0 {                                 // 만약 가리킬 수 있다면 코드를 수정해야 함
				prevLogTerm = ls.entries[prevIndex].Term
			}

			if in.GetPrevLogIndex() <= lastLogIndex && in.GetPrevLogTerm() == prevLogTerm {
				resp.Success = true

				logInsertIdx := in.GetPrevLogIndex() + 1
				newEntriesIdx := 0
				for {
					if logInsertIdx > lastLogIndex || newEntriesIdx >= len(in.GetEntries()) {
						break
					}
					if ls.entries[int(logInsertIdx)].Term != in.GetEntries()[newEntriesIdx].Term {
						break
					}
					logInsertIdx++
					newEntriesIdx++
				}

				if newEntriesIdx < len(in.GetEntries()) {
					ls.entries = append(ls.entries[:int(logInsertIdx)], in.GetEntries()[newEntriesIdx:]...)
				}

				if in.GetLeaderCommit() > ls.commitIndex {
					// TODO : 커밋 처리
				}
			}
		})

	}

	return resp, nil
}

func (n *RaftyNode) InstallSnapshot(in pb.Rafty_InstallSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}

func (n *RaftyNode) Write(ctx context.Context, in *pb.WriteReq) (*pb.WriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func (n *RaftyNode) Read(ctx context.Context, in *pb.ReadReq) (*pb.ReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (n *RaftyNode) FindLeader(ctx context.Context, in *pb.FindLeaderReq) (*pb.FindLeaderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLeader not implemented")
}
