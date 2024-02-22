package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webinar/microservice/auth-service/api"
	"webinar/microservice/auth-service/db/models"
	"webinar/microservice/auth-service/utils"
)

func (c *Client) Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var body api.AuthRequest
	err := decoder.Decode(&body)
	if err != nil {
		utils.ServeError(w, fmt.Errorf("decode err: %w", err), http.StatusBadRequest)
		return
	}
	user, err := c.user.GetByEmail(body.Email)
	if err != nil {
		utils.ServeError(w, err, http.StatusBadRequest)
		return
	}
	if !utils.CheckPassword(body.Password, user.Password) {
		utils.ServeError(w, err, http.StatusBadRequest)
		return
	}
	token, ttl, err := c.jwt.GenerateToken(user.Email)
	if err != nil {
		utils.ServeError(w, err, http.StatusInternalServerError)
		return
	}

	utils.ServeResponse(w, api.AuthResponse{
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		Token:     token,
		Ttl:       ttl,
	})
}

func (c *Client) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var body api.RegisterRequest
	err := decoder.Decode(&body)
	if err != nil {
		panic(err)
	}
	err = c.user.Create(&models.User{
		Email:    body.Email,
		Password: utils.HashPassword(body.Password),
	})
	if err != nil {
		utils.ServeError(w, err, http.StatusBadRequest)
		return
	}
	utils.ServeResponse(w, api.RegisterResponse{
		Status: "success",
	})
}
