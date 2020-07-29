package echo

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/tesarwijaya/ichos/config/database"
)

type (
	// Handler to hold collector
	Handler struct {
		collector *database.Collector
	}

	res struct {
		Messages database.MessageColletion
	}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Ws method to handle websocket connection
func (h *Handler) Ws(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		h.collector.Post(database.Message(string(msg)))
		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		res := res{
			Messages: h.collector.Get(),
		}

		json, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error Marshal: %s", err)
			return
		}

		// Write message back to browser
		if err = conn.WriteMessage(msgType, json); err != nil {
			return
		}
	}
}

// Init to initialize Websocket handler
func Init(collector *database.Collector) *Handler {
	return &Handler{
		collector,
	}
}
