package user

import "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"

type userRepo struct {
}

// Create implements UserRepository.
func (userRepo) Create(u models.User) (int, error) {
	panic("unimplemented")
}

// FindById implements UserRepository.
func (u userRepo) FindById(id uint) models.User {
	panic("unimplemented")
}

// IsRegistered implements UserRepository.
func (u userRepo) IsRegistered(login string) bool {
	panic("unimplemented")
}

func NewUserRepo(connection string) UserRepository {
	return userRepo{}
}
