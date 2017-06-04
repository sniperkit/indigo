package grpc

import (
	"fmt"
	"github.com/blevesearch/bleve/mapping"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type indigoGRPCServer struct {
	server   *grpc.Server
	listener net.Listener
	service  *service.IndigoGRPCService
}

func NewIndigoGRPCServer(port int, path string, indexMapping mapping.IndexMapping, indexType string, kvstore string, kvconfig map[string]interface{}) *indigoGRPCServer {
	server := grpc.NewServer()
	service := service.NewIndigoGRPCService(path, indexMapping, indexType, kvstore, kvconfig)

	proto.RegisterIndigoServer(server, service)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		log.WithFields(log.Fields{
			"port": port,
		}).Info("create listener")
	} else {
		log.WithFields(log.Fields{
			"port":  port,
			"error": err.Error(),
		}).Error("failed to create listener")
		return nil
	}

	return &indigoGRPCServer{
		server:   server,
		listener: listener,
		service:  service,
	}
}

func (igs *indigoGRPCServer) Start(deleteIndex bool) error {
	go func() {
		igs.service.OpenIndex(deleteIndex)
		igs.server.Serve(igs.listener)
		return
	}()

	log.WithFields(log.Fields{
		"addr": igs.listener.Addr().String(),
	}).Info("The Indigo gRPC Server started")

	return nil
}

func (igs *indigoGRPCServer) Stop(deleteIndex bool) error {
	igs.service.CloseIndex(deleteIndex)
	igs.server.GracefulStop()

	log.WithFields(log.Fields{
		"addr": igs.listener.Addr().String(),
	}).Info("The Indigo gRPC Server stopped")

	return nil
}
