package grpc

import (
	"fmt"
	"github.com/mosuka/indigo/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type indigoGRPCServer struct {
	server   *grpc.Server
	listener net.Listener
}

func NewIndigoGRPCServer(serverName string, serverPort int, indexDir string, indexMapping string, indexType string, indexStore string) *indigoGRPCServer {
	server := grpc.NewServer()

	proto.RegisterIndigoServer(server, NewIndigoGRPCService(indexDir, indexMapping, indexType, indexStore))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return nil
	}

	log.Printf("info: create Indigo gRPC server name=%s port=%d\n", serverName, serverPort)

	return &indigoGRPCServer{
		server:   server,
		listener: listener,
	}
}

func (brs *indigoGRPCServer) Start() error {
	go func() {
		brs.server.Serve(brs.listener)
		return
	}()

	log.Printf("info: Indigo gRPC server started\n")

	return nil
}

func (brs *indigoGRPCServer) Stop() error {
	brs.server.GracefulStop()

	log.Printf("info: Indigo gRPC server stopped\n")

	return nil
}
