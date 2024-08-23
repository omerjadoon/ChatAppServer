package handler

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mukeshkuiry/anycall/group"
	"github.com/mukeshkuiry/anycall/peer"
)

// Handler is the exported function Vercel will use to handle requests
func Handler(w http.ResponseWriter, r *http.Request) {
	router := mux.NewRouter()  // Use a different variable name for the router
	router.HandleFunc("/peer", peer.HandlePeerConnection)
	router.HandleFunc("/group", group.HandleGroupConnection)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Hello, World!"))
	}).Methods("GET")

	router.ServeHTTP(w, r)  // Serve HTTP using the router, not the request
}

func main() {
	log.Println("Server is running...")
}
