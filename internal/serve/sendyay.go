package serve

import (
	"net/http"

	"github.com/razvvan/CryptoYay-Server/internal/yay"
)

func SendYayHandler(w http.ResponseWriter, r *http.Request) {
	err := yay.Send("to", "org")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"status\": \"failed\"}"))
	}
}
