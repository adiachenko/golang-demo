package actions

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	message := Message{"Hello World"}

	if err := json.NewEncoder(w).Encode(message); err != nil {
		panic(err)
	}
}
