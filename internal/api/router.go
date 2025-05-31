package api

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/vm", CreateVMHandler).Methods("POST")
	r.HandleFunc("/vm", ListVMsHandler).Methods("GET")
	return r
}
