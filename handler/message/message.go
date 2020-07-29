package message

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tesarwijaya/ichos/config/database"
)

type (
	// Handler to hold collector
	Handler struct {
		collector *database.Collector
	}

	getRes struct {
		Status   string
		Messages database.MessageColletion
	}

	postBody struct {
		Message string
	}
	postRes struct {
		Status string
	}
)

// Post method to handle message post request
func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var b postBody
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.collector.Post(database.Message(b.Message))

	res := postRes{
		Status: "Ok",
	}

	json, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error Marshal: %s", err)
		return
	}

	w.Write(json)
}

// Get method to handle message get request
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	res := getRes{
		Status:   "Ok",
		Messages: h.collector.Get(),
	}

	json, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error Marshal: %s", err)
		return
	}

	w.Write(json)
}

// Init method to handle message handler initialization
func Init(collector *database.Collector) *Handler {
	return &Handler{
		collector,
	}
}
