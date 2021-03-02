package raft

type Config struct {
	ID     int64
	Listen string
	Peers  string // "1=localhost:9919,2=localhost:9920,3=localhost:9921"
}
