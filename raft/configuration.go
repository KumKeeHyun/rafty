package raft

import (
	"strconv"
	"strings"

	pb "github.com/KumKeeHyun/rafty/protobuf"
	"google.golang.org/grpc"
)

func newClusterConfig(cfg *Config) (*clusterConfig, error) {
	peers := make(map[int64]*peer, 0)
	ps := strings.Split(cfg.Peers, ",")
	for _, p := range ps {
		pInfo := strings.Split(p, "=")
		id, err := strconv.Atoi(pInfo[0])
		if err != nil {
			return nil, err
		}
		peer, err := newPeer(int64(id), pInfo[1])
		peers[int64(id)] = peer
	}

	return &clusterConfig{
		id:     cfg.ID,
		listen: cfg.Listen,
		peers:  peers,
	}, nil
}

type clusterConfig struct {
	id     int64
	listen string

	peers map[int64]*peer
}

func newPeer(id int64, addr string) (*peer, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &peer{
		id:   id,
		addr: addr,
		cli:  pb.NewRaftyClient(conn),
	}, nil
}

type peer struct {
	id   int64
	addr string
	cli  pb.RaftyClient
}
