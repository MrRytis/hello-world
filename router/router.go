package router

import (
	"github.com/MrRytis/hello-world/handler"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.Index).Methods("GET")
	r.HandleFunc("/api/hello", handler.HelloWorld).Methods("GET")

	return r
}
