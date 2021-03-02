package raft

import (
	"context"
	"sync"
)

const (
	stateLeader = iota
	stateCandidate
	stateFollower
)

func newNodeCtx() *nodeCtx {
	return &nodeCtx{
		state:         stateFollower,
		term:          1,
		votedFor:      -1,
		currentLeader: -1,
		opRequest:     make(chan nodeCtxReq),
	}
}

type nodeCtx struct {
	state         int
	term          uint64
	votedFor      int64
	currentLeader int64

	opRequest chan nodeCtxReq
}

type nodeCtxOp func(*nodeCtx)

type nodeCtxReq struct {
	op   nodeCtxOp
	done chan struct{}
}

func (n *nodeCtx) startNodeCtx(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case req := <-n.opRequest:
				req.op(n)
				req.done <- struct{}{}

			case <-ctx.Done():
				return
			}
		}
	}()
}

func (n *nodeCtx) Do(op nodeCtxOp) {
	done := make(chan struct{})
	n.opRequest <- nodeCtxReq{
		op:   op,
		done: done,
	}
	<-done
	close(done)
}
