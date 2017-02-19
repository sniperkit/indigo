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

	target := fmt.Sprintf("%s:%d", gRPCServerName, gRPCServerPort)

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err == nil {
		log.Printf("info: create connection target=\"%s\"\n", target)
	} else {
		log.Printf("error: failed to create connection (%s) target=\"%s\"\n", err.Error(), target)
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

	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", serverPath), handler.NewPutDocumentHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", serverPath), handler.NewGetDocumentHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", serverPath), handler.NewDeleteDocumentHandler(client)).Methods("DELETE")

	router.Handle(fmt.Sprintf("%s/{indexName}/_bulk", serverPath), handler.NewBulkHandler(client)).Methods("POST")

	router.Handle(fmt.Sprintf("%s/{indexName}/_search", serverPath), handler.NewSearchHandler(client)).Methods("POST")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err == nil {
		log.Printf("info: create listener port=%d\n", serverPort)
	} else {
		log.Printf("error: failed to create listener (%s) port=%d\n", err.Error(), serverPort)
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

	log.Printf("info: The Indigo REST Server started addr=\"%s\"\n", irs.listener.Addr().String())

	return nil
}

func (irs *indigoRESTServer) Stop() error {
	err := irs.conn.Close()
	if err == nil {
		log.Print("info: close connection\n")
	} else {
		log.Printf("error: failed to close connection (%s)\n", err.Error())
		return err
	}

	err = irs.listener.Close()
	if err == nil {
		log.Print("info: close listener\n")
	} else {
		log.Printf("error: failed to close listener (%s)\n", err.Error())
		return err
	}

	log.Printf("info: The Indigo REST Server stopped addr=\"%s\"\n", irs.listener.Addr().String())

	return nil
}
