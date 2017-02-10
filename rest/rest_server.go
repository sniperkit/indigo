package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net"
	"net/http"
)

type indigoRESTServer struct {
	router   *mux.Router
	listener net.Listener
}

func NewIndigoRESTServer(serverName string, serverPort int) *indigoRESTServer {

	router := mux.NewRouter()
	router.StrictSlash(true)

	router.Handle("/api/{indexName}", NewCreateIndexHandler("aaa")).Methods("GET")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return nil
	}

	log.Printf("info: create Indigo REST server name=%s port=%d\n", serverName, serverPort)

	return &indigoRESTServer{
		router:   router,
		listener: listener,
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
	err := brs.listener.Close()
	if err != nil {
		log.Printf("error: %s\n", err.Error())
		return err
	}

	log.Printf("info: Indigo REST server stopped\n")

	return nil
}
