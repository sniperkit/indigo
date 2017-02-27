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
	service  *indigoGRPCService
}

func NewIndigoGRPCServer(port int, dataDir string) *indigoGRPCServer {
	server := grpc.NewServer()
	service := NewIndigoGRPCService(dataDir)

	proto.RegisterIndigoServer(server, service)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		log.Printf("info: create listener port=%d \n", port)
	} else {
		log.Printf("error: %s port=%d \n", err.Error(), port)
		return nil
	}

	return &indigoGRPCServer{
		server:   server,
		listener: listener,
		service:  service,
	}
}

func (igs *indigoGRPCServer) Start(openExistsIndex bool) error {
	go func() {
		if openExistsIndex {
			igs.service.OpenIndices()
		}
		igs.server.Serve(igs.listener)
		return
	}()

	log.Printf("info: The Indigo gRPC Server started addr=\"%s\"\n", igs.listener.Addr().String())

	return nil
}

func (igs *indigoGRPCServer) Stop() error {
	igs.service.CloseIndices()
	igs.server.GracefulStop()

	log.Printf("info: The Indigo gRPC Server stopped addr=\"%s\"\n", igs.listener.Addr().String())

	return nil
}
