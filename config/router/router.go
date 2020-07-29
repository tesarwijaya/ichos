package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tesarwijaya/ichos/config/database"
	"github.com/tesarwijaya/ichos/handler/echo"
	"github.com/tesarwijaya/ichos/handler/message"
)

// Init router config
func Init(col *database.Collector) *mux.Router {
	r := mux.NewRouter()

	messageHandler := message.Init(col)
	r.HandleFunc("/message", messageHandler.Get).Methods("GET")
	r.HandleFunc("/message", messageHandler.Post).Methods("POST")

	echoHandler := echo.Init(col)
	r.HandleFunc("/echo", echoHandler.Ws)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})
	return r
}
