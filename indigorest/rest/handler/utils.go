package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, error string, code int) {

	msg := struct {
		error string `json:"error"`
	}{
		error: error,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		log.Printf("warn: %s\n", err.Error())
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	w.Write(b)
}
