package user

import (
	"errors"
	"fmt"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
)

type UserRepository interface {
	Create(u models.User) (int, error)
	IsRegistered(login string) bool
	Auth(u models.User) (int, bool)
	FindById(id uint) (*models.User, error)
}

type user struct {
	repo UserRepository
}

// Auth implements handlers.UserManager.
func (u user) Auth(user models.User) (*models.UserAuth, error) {
	id, isFound := u.repo.Auth(user)

	if !isFound {
		return nil, errors.New("401")
	}

	return &models.UserAuth{
		//TODO: JWT
		Jwt: fmt.Sprint(id),
	}, nil
}

// Register implements handlers.UserManager.
func (u user) Register(user models.User) (*models.UserAuth, error) {
	if u.repo.IsRegistered(user.Login) {
		return nil, errors.New("already registered")
	}

	id, err := u.repo.Create(user)

	if err != nil {
		return nil, err
	}

	return &models.UserAuth{
		//TODO: JWT
		Jwt: fmt.Sprint(id),
	}, nil
}

func NewUser(repo UserRepository) user {
	return user{
		repo: repo,
	}
}
