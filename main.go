package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

var r *mux.Router

func init() {
	staticDir := "/clientApp/public"
	r = mux.NewRouter()
	r.Handle("/ws", websocket.Handler(HandleNewConnection))
	r.HandleFunc("/config", saveConfig).Methods("POST")
	r.HandleFunc("/config", getConfigs).Methods("GET")
	r.HandleFunc("/config/{id}", getConfig).Methods("GET")
	r.PathPrefix("/").
		Handler(http.StripPrefix("/", http.FileServer(http.Dir("."+staticDir))))
	// r.Handle("/", http.FileServer(http.Dir("./clientApp/public")))
}

func main() {
	log.Fatal(http.ListenAndServe(":3000", r))
}