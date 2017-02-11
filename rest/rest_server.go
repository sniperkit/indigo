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

func NewIndigoRESTServer(serverName string, serverPort int, gRPCServerName string, gRPCServerPort int) *indigoRESTServer {

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
	router.Handle("/api/mapping", handler.NewGetMappingHandler(client)).Methods("GET")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return nil
	}

	log.Printf("info: create Indigo REST server name=%s port=%d\n", serverName, serverPort)

	return &indigoRESTServer{
		router:   router,
		listener: listener,
		conn:     conn,
	}
}

func (brs *indigoRESTServer) Start() error {
	go func() {
		http.Serve(brs.listener, brs.router)
		return
	}()

	log.Printf("info: Indigo REST server started\n")

	return nil
}

func (brs *indigoRESTServer) Stop() error {
	brs.conn.Close()

	err := brs.listener.Close()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	log.Printf("info: Indigo REST server stopped\n")

	return nil
}
