package group

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/mukeshkuiry/anycall/models"
)

func HandleGroupConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := models.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// get room id from url
	roomID := r.URL.Query().Get("room_id")

	if roomID == "" {
		log.Println("Room ID is required")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Room ID is required"))
		return
	}

	// create room if not exists
	if models.GroupConnectionRooms[roomID] == nil {
		models.GroupConnectionRooms[roomID] = &models.Room{
			ID:      roomID,
			Clients: make(map[*websocket.Conn]bool),
		}
	}

	// add client to room
	models.GroupConnectionRooms[roomID].Clients[conn] = true

	log.Println("Client connected to room: ", roomID)

	// listen for messages
	for {
		var msg []byte
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message: ", err)
			delete(models.GroupConnectionRooms[roomID].Clients, conn)
			break
		}
		log.Println("Message received: ", string(msg))
		// broadcast message to all clients in room

		for client := range models.GroupConnectionRooms[roomID].Clients {
			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Println("Error writing message: ", err)
				delete(models.GroupConnectionRooms[roomID].Clients, client)
				break
			}
		}
	}
	// close connection when function returns

	defer conn.Close()
}

// design the group video call feature
