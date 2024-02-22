package api

import "time"

type AuthResponse struct {
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	Token     string    `json:"token"`
	Ttl       int64     `json:"ttl"`
}

type RegisterResponse struct {
	Status string
}
