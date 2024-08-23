package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukeshkuiry/anycall/group"
	"github.com/mukeshkuiry/anycall/peer"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/peer", peer.HandlePeerConnection)
	r.HandleFunc("/group", group.HandleGroupConnection)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		// allow cross origin requests
		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Write([]byte("Hello, World!"))
	}).Methods("GET")
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", r))

}
