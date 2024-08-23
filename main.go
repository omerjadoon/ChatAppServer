package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukeshkuiry/anycall/group"
	"github.com/mukeshkuiry/anycall/peer"
)

// Handler is the exported function Vercel will use to handle requests
func Handler(w http.ResponseWriter, r *http.Request) {
	r := mux.NewRouter()
	r.HandleFunc("/peer", peer.HandlePeerConnection)
	r.HandleFunc("/group", group.HandleGroupConnection)
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	r.ServeHTTP(w, r)
}

func main() {
	log.Println("Server is running...")
}
