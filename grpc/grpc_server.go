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

	log.Printf("info: The Indigo gRPC Server created name=%s port=%d\n", serverName, serverPort)

	return &indigoGRPCServer{
		server:   server,
		listener: listener,
	}
}

func (igs *indigoGRPCServer) Start() error {
	go func() {
		igs.server.Serve(igs.listener)
		log.Print("info: The Indigo gRPC Server started\n")
		return
	}()

	return nil
}

func (igs *indigoGRPCServer) Stop() error {
	igs.server.GracefulStop()
	log.Print("info: The Indigo gRPC Server stopped\n")

	return nil
}
