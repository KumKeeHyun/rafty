package raft

import (
	"context"
	"sync"

	pb "github.com/KumKeeHyun/rafty/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRaftyNode(cfg *Config) (*raftyNode, error) {
	clusterConfig, err := newClusterConfig(cfg)
	if err != nil {
		return nil, err
	}

	wg := &sync.WaitGroup{}
	nodeCtx, nodeCancel := context.WithCancel(context.Background())

	nodeContext := newNodeCtx()
	logStorage := newLogStorage()

	nodeContext.startNodeCtx(nodeCtx, wg)
	logStorage.startLogStorage(nodeCtx, wg)

	return &raftyNode{
		config:                   clusterConfig,
		nodeContext:              nodeContext,
		logStorage:               logStorage,
		wg:                       wg,
		ctxCancel:                nodeCancel,
		UnimplementedRaftyServer: pb.UnimplementedRaftyServer{},
	}, nil
}

type raftyNode struct {
	config      *clusterConfig
	nodeContext *nodeCtx    // term, votedFor, state
	logStorage  *logStorage // log entries, snapshot

	wg        *sync.WaitGroup    // raftyNode 하위에 실행된 고루틴 담당
	ctxCancel context.CancelFunc // Candidate, Leader 상태에서 실행한 고루틴을 중지시키는 함수. becomeFollower 에서 호출

	pb.UnimplementedRaftyServer
}

func (n *raftyNode) startElection() {

}

func (n *raftyNode) startLeader() {

}

func (n *raftyNode) sendHeartbeats() {

}

func (n *raftyNode) becomeFollower() {

}

// ============= grpc ==============

func (n *raftyNode) RequestVote(ctx context.Context, in *pb.RequestVoteReq) (*pb.RequestVoteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestVote not implemented")
}

func (n *raftyNode) AppendEntries(ctx context.Context, in *pb.AppendEntriesReq) (*pb.AppendEntriesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}

func (n *raftyNode) InstallSnapshot(in pb.Rafty_InstallSnapshotServer) error {
	return status.Errorf(codes.Unimplemented, "method InstallSnapshot not implemented")
}

func (n *raftyNode) Write(ctx context.Context, in *pb.WriteReq) (*pb.WriteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func (n *raftyNode) Read(ctx context.Context, in *pb.ReadReq) (*pb.ReadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (n *raftyNode) FindLeader(ctx context.Context, in *pb.FindLeaderReq) (*pb.FindLeaderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindLeader not implemented")
}
