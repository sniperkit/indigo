package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/bleve-server/http_handler"
	"log"
	"net"
	"net/http"
)

type bleveRESTServer struct {
	router   *mux.Router
	listener net.Listener
}

func NewBleveRESTServer(serverName string, serverPort int) *bleveRESTServer {

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.Handle("/api/{indexName}", http_handler.NewCreateIndexHandler("aaa")).Methods("GET")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return nil
	}

	log.Printf("info: create Bleve REST server name=%s port=%d\n", serverName, serverPort)

	return &bleveRESTServer{
		router:   router,
		listener: listener,
	}
}

func (brs *bleveRESTServer) Start() error {
	go func() {
		http.Serve(brs.listener, brs.router)
		return
	}()

	log.Printf("info: Bleve REST server started\n")

	return nil
}

func (brs *bleveRESTServer) Stop() error {
	err := brs.listener.Close()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	log.Printf("info: Bleve REST server stopped\n")

	return nil
}
