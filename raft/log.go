package raft

import pb "github.com/KumKeeHyun/rafty/protobuf"

type logStorage struct {
	entries []*pb.Entry
}
