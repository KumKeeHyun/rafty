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
	nextIndex := make(map[int64]uint64, len(n.config.peers))
	matchIndex := make(map[int64]uint64, len(n.config.peers))
	for peerID := range n.config.peers {
		nextIndex[peerID] = lastLogIndex + 1
		matchIndex[peerID] = 0
	}

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
				hbCtx, hbCancel = context.WithTimeout(context.Background(), sendHeartbeatsDuration)
				n.sendHeartbeats(hbCtx)

			case <-leaderCtx.Done():
				hbCancel()
				return
			}
		}
	}()
}

func (n *RaftyNode) sendHeartbeats(ctx context.Context) {

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
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
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
