package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/proto"
	"github.com/mosuka/indigo/rest/handler"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type indigoRESTServer struct {
	router   *mux.Router
	listener net.Listener
	conn     *grpc.ClientConn
}

func NewIndigoRESTServer(serverPort int, serverPath, gRPCServerName string, gRPCServerPort int) *indigoRESTServer {

	router := mux.NewRouter()
	router.StrictSlash(true)

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort), grpc.WithInsecure())
	if err != nil {
		return nil
	}

	client := proto.NewIndigoClient(conn)

	/*
	 * set handlers
	 */
	router.Handle(fmt.Sprintf("%s/{indexName}", serverPath), handler.NewCreateIndexHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}", serverPath), handler.NewDeleteIndexHandler(client)).Methods("DELETE")

	router.Handle(fmt.Sprintf("%s/{indexName}/_mapping", serverPath), handler.NewGetMappingHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/_stats", serverPath), handler.NewGetStatsHandler(client)).Methods("GET")

	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", serverPath), handler.NewIndexDocumentHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", serverPath), handler.NewDeleteDocumentHandler(client)).Methods("DELETE")

	router.Handle(fmt.Sprintf("%s/{indexName}/_bulk", serverPath), handler.NewIndexBulkHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_bulk", serverPath), handler.NewDeleteBulkHandler(client)).Methods("DELETE")

	router.Handle(fmt.Sprintf("%s/{indexName}/_search", serverPath), handler.NewSearchHandler(client)).Methods("POST")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: failed to create listener :%d (%s)\n", serverPort, err.Error())
		return nil
	}

	log.Printf("info: The Indigo REST Server created %s\n", listener.Addr().String())

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

	log.Printf("info: The Indigo REST Server started %s\n", irs.listener.Addr().String())

	return nil
}

func (irs *indigoRESTServer) Stop() error {
	err := irs.conn.Close()
	if err != nil {
		log.Printf("error: failed to close connection (%s)\n", err.Error())
		return err
	}

	err = irs.listener.Close()
	if err != nil {
		log.Printf("error: failed to close listener (%s)\n", err.Error())
		return err
	}

	log.Printf("info: The Indigo REST Server stopped %s\n", irs.listener.Addr().String())

	return nil
}
