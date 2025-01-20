package user

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type UserRepository interface {
	Create(u models.User) (int, error)
	IsRegistered(login string) bool
	FindById(id uint) models.User
}

type user struct {
	repo UserRepository
}

// Auth implements handlers.UserManager.
func (user) Auth(u models.User) (models.UserAuth, error) {
	panic("unimplemented")
}

// Register implements handlers.UserManager.
func (user) Register(u models.User) (models.UserAuth, error) {
	panic("unimplemented")
}

func NewUser(repo UserRepository) user {
	return user{
		repo: repo,
	}
}
