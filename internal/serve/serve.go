package serve

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func Listen() error {
	r := mux.NewRouter()
	r.HandleFunc("/send_yay", SendYayHandler)

	listenAddr := viper.GetString("bind-address")
	if listenAddr == "" {
		listenAddr = "127.0.0.1:8000"
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         listenAddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return srv.ListenAndServe()
}
