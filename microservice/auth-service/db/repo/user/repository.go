package user

import (
	"errors"
	"sync"
	"time"
	"webinar/microservice/auth-service/db"
	"webinar/microservice/auth-service/db/models"
)

type User interface {
	GetByEmail(email string) (*models.User, error)
	Create(user *models.User) error
}

type user struct {
	db sync.Map
}

func (u *user) GetByEmail(email string) (*models.User, error) {
	val, ok := u.db.Load(email)
	if !ok {
		return nil, db.NotExists
	}
	return val.(*models.User), nil
}

func (u *user) Create(user *models.User) error {
	model, err := u.GetByEmail(user.Email)
	if err != nil && !errors.Is(err, db.NotExists) {
		return err
	}
	if model != nil {
		return db.AlreadyExists
	}
	user.CreatedAt = time.Now()
	u.db.Store(user.Email, user)
	return nil
}

func New() User {
	return &user{}
}
