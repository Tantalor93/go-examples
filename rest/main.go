package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/tantalor93/go-examples/rest/domain"
	"net/http"
)

func main() {
	log.Info("Starting server")

	router := mux.NewRouter()
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		json.NewEncoder(writer).Encode(domain.User{Id: "a1", Name: "Petr"})
		writer.WriteHeader(http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8081", router))
}
