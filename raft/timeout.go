package raft

import (
	"context"
	"math/rand"
	"sync"
	"time"
)

var (
	once  sync.Once
	timer *timeout
)

// getTimer 싱글톤 timeout
func getTimer() *timeout {
	once.Do(func() {
		timer = newTimeout(electionTimeoutDuration, minimumElectionTimeoutDuration)
	})
	return timer
}

const (
	timeoutRun = iota
	timeoutStop
)

type timeout struct {
	state     int
	election  time.Duration
	minElect  time.Duration
	timestamp time.Time
	resetReq  chan chan struct{}
	stopReq   chan chan struct{}
}

// candidate 선거 타이머에 따라 후보자 루틴을 수행할 객체
type candidate interface {
	startElection()
}

func newTimeout(elect, minElect time.Duration) *timeout {
	return &timeout{
		state:     timeoutStop,
		election:  elect,
		minElect:  minElect,
		timestamp: time.Time{},
		resetReq:  make(chan chan struct{}),
		stopReq:   make(chan chan struct{}),
	}
}

// electionTimeout 클러스터 노드들이 같은 타이머를 갖지 않도록 랜덤한 타이머 생성
func (t *timeout) randElectionTimeout() *time.Ticker {
	randTimeout := time.Duration(rand.Intn(50)) * time.Millisecond
	return time.NewTicker(randTimeout + t.election)
}

// startTimer 노드가 선거를 시작하도록 하는 타이머 시작
func (t *timeout) startTimer(ctx context.Context, wg *sync.WaitGroup, c candidate) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		timeout := t.randElectionTimeout()
		t.timestamp = time.Now()
		t.state = timeoutRun

		for {
			select {
			case <-timeout.C:
				// log election timeout, start election
				c.startElection()
				t.timestamp = time.Now()

			case done := <-t.resetReq:
				timeout.Stop()
				timeout = t.randElectionTimeout()
				t.timestamp = time.Now()
				t.state = timeoutRun
				done <- struct{}{}

			case done := <-t.stopReq:
				timeout.Stop()
				t.state = timeoutStop
				done <- struct{}{}

			case <-ctx.Done():
				return
			}
		}
	}()
}

// reset 새로운 선거를 위한 타이머를 재시작함
// 1. 다른 리더의 존재를 확인했을 때
// 2. AppendEntries 호출을 받았을 때
func (t *timeout) reset() {
	req := make(chan struct{})
	t.resetReq <- req
	<-req
	close(req)
}

// stop 새로운 선거를 위한 타이머를 멈춤
func (t *timeout) stop() {
	req := make(chan struct{})
	t.stopReq <- req
	<-req
	close(req)
}

// isValidTimeout 최소 선거 시작 시간인지 확인
// 최소 선거 시작 시간이 아닌 경우 RequestVote 호출을 무시함
func (t *timeout) isValidTimeout() bool {
	if t.state == timeoutRun {
		//
		validTimeout := t.timestamp.Add(t.election - t.minElect)
		return time.Now().After(validTimeout)
	}

	return false
}
