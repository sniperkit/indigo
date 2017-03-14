package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/rest/handler"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

type indigoRESTServer struct {
	router   *mux.Router
	listener net.Listener
	conn     *grpc.ClientConn
}

func NewIndigoRESTServer(port int, basePath, gRPCServer string) *indigoRESTServer {
	router := mux.NewRouter()
	router.StrictSlash(true)

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err == nil {
		log.WithFields(log.Fields{
			"grpc_server": gRPCServer,
		}).Info("succeeded in creating connection")
	} else {
		log.WithFields(log.Fields{
			"grpc_server": gRPCServer,
			"err":         err,
		}).Error("failed to create connection")

		return nil
	}

	client := proto.NewIndigoClient(conn)

	/*
	 * set handlers
	 */
	router.Handle(fmt.Sprintf("%s/_list", basePath), handler.NewListIndicesHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}", basePath), handler.NewCreateIndexHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}", basePath), handler.NewGetIndexHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}", basePath), handler.NewDeleteIndexHandler(client)).Methods("DELETE")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewPutDocumentHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewGetDocumentHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewDeleteDocumentHandler(client)).Methods("DELETE")
	router.Handle(fmt.Sprintf("%s/{indexName}/_open", basePath), handler.NewOpenIndexHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_close", basePath), handler.NewCloseIndexHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_bulk", basePath), handler.NewBulkHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_search", basePath), handler.NewSearchHandler(client)).Methods("POST")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		log.WithFields(log.Fields{
			"port": port,
		}).Info("succeeded in creating listener")
	} else {
		log.WithFields(log.Fields{
			"port": port,
			"err":  err,
		}).Error("failed to create listener")
		return nil
	}

	return &indigoRESTServer{
		router:   router,
		listener: listener,
		conn:     conn,
	}
}

func (irs *indigoRESTServer) Start() error {
	go func() {
		http.Serve(irs.listener, irs.router)
		return
	}()

	log.WithFields(log.Fields{
		"addr": irs.listener.Addr().String(),
	}).Info("The Indigo REST Server started")

	return nil
}

func (irs *indigoRESTServer) Stop() error {
	err := irs.conn.Close()
	if err == nil {
		log.Info("succeeded in closing connection")
	} else {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to close connection")

		return err
	}

	err = irs.listener.Close()
	if err == nil {
		log.Info("succeeded in closing listener")
	} else {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("failed to close listener")

		return err
	}

	log.WithFields(log.Fields{
		"addr": irs.listener.Addr().String(),
	}).Info("The Indigo REST Server stopped")

	return nil
}
