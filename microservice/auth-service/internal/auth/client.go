package auth

import (
	"webinar/microservice/auth-service/config"
	"webinar/microservice/auth-service/db/repo/user"
	"webinar/microservice/auth-service/pkg/token"
)

type Client struct {
	user user.User
	jwt  token.JWT
}

func NewClient(cfg config.Config) *Client {
	return &Client{
		user: user.New(),
		jwt:  token.New(cfg.JWT.Ttl, cfg.JWT.Issuer, cfg.JWT.Secret),
	}
}
