package user

import (
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	"github.com/Masterminds/squirrel"
)

type userRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
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

func NewUserRepo(db db.DB) UserRepository {
	return userRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
