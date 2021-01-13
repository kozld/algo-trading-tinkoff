package trader

import (
	"fmt"
	"tinkoff-trade-bot/trader/config"
	"log"
	"net"

	pb "tinkoff-trade-bot/trader/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GRPCServer struct {
	traderService *TraderService
	config      *config.TraderConfig
}

func NewGRPCServer(s *TraderService, cfg *config.TraderConfig) *GRPCServer {
	return &GRPCServer{
		traderService: s,
		config:      cfg,
	}
}

func (s *GRPCServer) Start() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.config.RPCListenPort))
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	rpcServer := grpc.NewServer()

	pb.RegisterTraderServer(rpcServer, pb.TraderServer(s.traderService))
	reflection.Register(rpcServer)

	err = rpcServer.Serve(listener)
	if err != nil {
		log.Fatal("failed to serve", err)
	}
}