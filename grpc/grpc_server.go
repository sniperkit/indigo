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

func NewIndigoGRPCServer(serverPort int, dataDir string) *indigoGRPCServer {
	server := grpc.NewServer()

	proto.RegisterIndigoServer(server, NewIndigoGRPCService(dataDir))
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: failed to create listener (%s) port=%d \n", err.Error(), serverPort)
		return nil
	}

	log.Printf("info: The Indigo gRPC Server created port=%d\n", serverPort)

	return &indigoGRPCServer{
		server:   server,
		listener: listener,
	}
}

func (igs *indigoGRPCServer) Start() error {
	go func() {
		igs.server.Serve(igs.listener)
		return
	}()

	log.Printf("info: The Indigo gRPC Server started addr=\"%s\"\n", igs.listener.Addr().String())

	return nil
}

func (igs *indigoGRPCServer) Stop() error {
	igs.server.GracefulStop()

	log.Printf("info: The Indigo gRPC Server stopped addr=\"%s\"\n", igs.listener.Addr().String())

	return nil
}
