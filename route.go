package main

import (
	"github.com/chunganhbk/chat-golang/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	//init router
	r := mux.NewRouter()
	var HubGlob = newHub()
	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(HubGlob, w, r)
	})
	api := r.PathPrefix("/v1/api").Subrouter()
	group := api.PathPrefix("/group").Subrouter()
	group.HandleFunc("/", controller.CreateGroupHandle).Methods("POST")
	return r
}