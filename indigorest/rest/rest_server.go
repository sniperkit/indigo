//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mosuka/indigo/indigorest/rest/handler"
	"github.com/mosuka/indigo/proto"
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

func NewIndigoRESTServer(port int, basePath, server string) *indigoRESTServer {
	router := mux.NewRouter()
	router.StrictSlash(true)

	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err == nil {
		log.WithFields(log.Fields{
			"server": server,
		}).Info("succeeded in creating connection")
	} else {
		log.WithFields(log.Fields{
			"server": server,
			"err":    err,
		}).Error("failed to create connection")

		return nil
	}

	client := proto.NewIndigoClient(conn)

	/*
	 * set handlers
	 */
	router.Handle(fmt.Sprintf("%s/", basePath), handler.NewGetIndexHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{id}", basePath), handler.NewPutDocumentHandler(client)).Methods("PUT")
	router.Handle(fmt.Sprintf("%s/{id}", basePath), handler.NewGetDocumentHandler(client)).Methods("GET")
	router.Handle(fmt.Sprintf("%s/{id}", basePath), handler.NewDeleteDocumentHandler(client)).Methods("DELETE")
	router.Handle(fmt.Sprintf("%s/_bulk", basePath), handler.NewBulkHandler(client)).Methods("POST")
	router.Handle(fmt.Sprintf("%s/_search", basePath), handler.NewSearchHandler(client)).Methods("POST")

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
