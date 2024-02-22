package main

import (
	"net/http"
	"webinar/microservice/auth-service/config"
	"webinar/microservice/auth-service/internal/auth"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	client := auth.NewClient(*cfg)
	mux.HandleFunc("/login", client.Auth)
	mux.HandleFunc("/register", client.Register)

	if err = http.ListenAndServe(cfg.Port, mux); err != nil {
		panic(err)
	}
}
