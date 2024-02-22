package utils

import (
	"encoding/json"
	"net/http"
	"webinar/microservice/auth-service/api"
)

func ServeError(w http.ResponseWriter, err error, code int) {
	jsonResp, _ := json.Marshal(api.ErrResponse{
		Error: err.Error(),
	})
	_, _ = w.Write(jsonResp)
	w.WriteHeader(code)
}

func ServeResponse(w http.ResponseWriter, data any) {
	jsonResp, _ := json.Marshal(data)
	_, _ = w.Write(jsonResp)
	w.WriteHeader(http.StatusOK)
}
