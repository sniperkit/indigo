package http_handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CreateIndexHandler struct {
	basePath        string
	IndexNameLookup func(req *http.Request) string
}

func NewCreateIndexHandler(basePath string) *CreateIndexHandler {
	return &CreateIndexHandler{
		basePath: basePath,
	}
}

func (h *CreateIndexHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	log.Println("Responsing to /hello request")
	log.Println(req.UserAgent())

	vars := mux.Vars(req)
	name := vars["indexName"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Hello:", name)

}
