package task

import (
	"webinar/microservice/task-service/config"
	"webinar/microservice/task-service/pkg/token"
)

type Client struct {
	jwt token.JWT
}

func NewClient(cfg config.Config) *Client {
	return &Client{
		jwt: token.New(cfg.JWT.Ttl, cfg.JWT.Issuer, cfg.JWT.Secret),
	}
}
