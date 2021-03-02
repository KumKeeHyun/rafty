package raft

import "time"

type Config struct {
	ID     int64
	Listen string
	Peers  string // "1=localhost:9919,2=localhost:9920,3=localhost:9921"
}

var (
	electionTimeoutDuration        = 300 * time.Millisecond
	minimumElectionTimeoutDuration = 15 * time.Millisecond
	sendHeartbeatsDuration         = 50 * time.Millisecond
)
