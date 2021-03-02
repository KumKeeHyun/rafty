package raft

import (
	"context"
	"sync"

	pb "github.com/KumKeeHyun/rafty/protobuf"
)

func newLogStorage() *logStorage {
	entries := make([]*pb.Entry, 0, 1000)
	entries = append(entries, &pb.Entry{
		Term:      1,
		Index:     0,
		RequestID: 0,
		Type:      pb.EntryType_NoOperation,
		Data:      []byte{},
	})
	return &logStorage{
		entries:          entries,
		snapshotMetadata: nil,
		opRequest:        make(chan logRequest),
	}
}

type logStorage struct {
	entries          []*pb.Entry // 길이는 항상 1 이상으로 유지해야 함
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

// lastLogIndexAndTerm 마지막 로그 엔트리의 index, term 반환
// 동시성을 위해 단독으로 호출하면 안됨. 항상 Do 함수 안에서 호출
func (l *logStorage) lastLogIndexAndTerm() (index, term uint64) {
	lastLog := l.entries[len(l.entries)-1] // log 의 길이는 항상 1 이상으로 유지
	index = lastLog.Index
	term = lastLog.Term
	return
}

// appendEntry Entry.Index를 알아서 입력해준 후 entries에 추가, 추가된 entry의 index 반환
// 동시성을 위해 단독으로 호출하면 안됨. 항상 Do 함수 안에서 호출
func (l *logStorage) appendEntry(entry *pb.Entry) uint64 {
	entry.Index = l.entries[len(l.entries)-1].Index + 1
	l.entries = append(l.entries, entry)
	return entry.Index
}
