package raft

import (
	"context"
	"sync"

	pb "github.com/KumKeeHyun/rafty/protobuf"
)

func newLogStorage() *logStorage {
	return &logStorage{
		entries:          make([]*pb.Entry, 1000),
		snapshotMetadata: nil,
		opRequest:        make(chan logRequest),
	}
}

type logStorage struct {
	entries          []*pb.Entry
	snapshotMetadata *snapshot

	opRequest chan logRequest
}

type snapshot struct {
	fileName     string
	lastLogIndex uint64
	lastLogTerm  uint64
}

type logOp func(*logStorage)

type logRequest struct {
	op   logOp
	done chan struct{}
}

func (l *logStorage) startLogStorage(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case req := <-l.opRequest:
				req.op(l)
				req.done <- struct{}{}

			case <-ctx.Done():
				return
			}
		}
	}()
}

func (l *logStorage) Do(op logOp) {
	done := make(chan struct{})
	l.opRequest <- logRequest{
		op:   op,
		done: done,
	}
	<-done
	close(done)
}
