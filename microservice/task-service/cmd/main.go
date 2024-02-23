package main

import (
	"net/http"
	"webinar/microservice/task-service/config"
	"webinar/microservice/task-service/internal/task"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	client := task.NewClient(*cfg)
	mux.HandleFunc("/get", client.Get)

	if err = http.ListenAndServe(cfg.Port, mux); err != nil {
		panic(err)
	}
}
