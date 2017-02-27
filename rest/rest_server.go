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

func NewIndigoRESTServer(port int, basePath, gRPCServer string) *indigoRESTServer {
	router := mux.NewRouter()
	router.StrictSlash(true)

	conn, err := grpc.Dial(gRPCServer, grpc.WithInsecure())
	if err == nil {
		log.Printf("info: create connection target=\"%s\"\n", gRPCServer)
	} else {
		log.Printf("error: %s target=\"%s\"\n", err.Error(), gRPCServer)
		return nil
	}

	client := proto.NewIndigoClient(conn)

	/*
	 * set handlers
	 */
	router.Handle(fmt.Sprintf("%s/{indexName}", basePath), handler.NewCreateIndexHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}", basePath), handler.NewDeleteIndexHandler(client)).Methods("DELETE")
	router.Handle(fmt.Sprintf("%s/{indexName}/_open", basePath), handler.NewOpenIndexHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_close", basePath), handler.NewCloseIndexHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_mapping", basePath), handler.NewGetMappingHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/_stats", basePath), handler.NewGetStatsHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewPutDocumentHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewGetDocumentHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{indexName}/{id}", basePath), handler.NewDeleteDocumentHandler(client)).Methods("DELETE")
	router.Handle(fmt.Sprintf("%s/{indexName}/_bulk", basePath), handler.NewBulkHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/{indexName}/_search", basePath), handler.NewSearchHandler(client)).Methods("POST")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err == nil {
		log.Printf("info: create listener port=%d\n", port)
	} else {
		log.Printf("error: %s port=%d\n", err.Error(), port)
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
		log.Printf("error: %s\n", err.Error())
		return err
	}

	err = irs.listener.Close()
	if err == nil {
		log.Print("info: close listener\n")
	} else {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	log.Printf("info: The Indigo REST Server stopped addr=\"%s\"\n", irs.listener.Addr().String())

	return nil
}
