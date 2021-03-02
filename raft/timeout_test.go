package raft

import (
	"sync"
	"testing"
	"time"

	"golang.org/x/net/context"
)

type mockCandidate struct {
	out chan struct{}
}

func (mc *mockCandidate) startElection() {
	mc.out <- struct{}{}
}

func TestReset(t *testing.T) {
	elect := 500 * time.Millisecond
	minElect := 100 * time.Millisecond
	timeout := newTimeout(elect, minElect)

	out := make(chan struct{})
	mc := mockCandidate{out: out}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	timeout.startTimer(ctx, &wg, &mc)

	time.Sleep(200 * time.Millisecond)
	timeout.reset()

	select {
	case <-out:
		cancel()
		wg.Wait()
		t.Fail()
	case <-time.After(400 * time.Millisecond):
		cancel()
		wg.Wait()
		// 200ms가 지나서 reset 했기 때문에 200ms+500ms 이내에 startElection이 호출되지 않음
	}
}

func TestStop(t *testing.T) {
	elect := 500 * time.Millisecond
	minElect := 100 * time.Millisecond
	timeout := newTimeout(elect, minElect)

	out := make(chan struct{})
	mc := mockCandidate{out: out}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	timeout.startTimer(ctx, &wg, &mc)

	timeout.stop()

	select {
	case <-out:
		cancel()
		wg.Wait()
		t.Fail()
	case <-time.After(600 * time.Millisecond):
		cancel()
		wg.Wait()
		// 실제 타임아웃이 지나도 startElection이 호출되지 않음
	}
}

func TestIsValidTimeout(t *testing.T) {
	elect := 500 * time.Millisecond
	minElect := 100 * time.Millisecond
	timeout := newTimeout(elect, minElect)

	out := make(chan struct{})
	mc := mockCandidate{out: out}
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	timeout.startTimer(ctx, &wg, &mc)

	time.Sleep(450 * time.Millisecond)
	if !timeout.isValidTimeout() {
		cancel()
		wg.Wait()
		t.Fail()
	}
	cancel()
	wg.Wait()
}
