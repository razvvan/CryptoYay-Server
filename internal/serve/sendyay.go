package serve

import (
	"encoding/json"
	"net/http"

	"github.com/razvvan/CryptoYay-Server/internal/yay"
)

type SendYayRequest struct {
	To  string `json:"to"`
	Org string `json:"org"`
}

func SendYayHandler(w http.ResponseWriter, r *http.Request) {
	syr := SendYayRequest{}

	err := json.NewDecoder(r.Body).Decode(&syr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = yay.Send(syr.To, syr.Org)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
