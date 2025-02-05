package user

import (
	"github.com/sirupsen/logrus"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	servererrors "github.com/LexusEgorov/go-musthave-diploma-tpl/internal/serverErrors"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/utils"
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
		err := servererrors.UnauthorizedError{Message: "unknown credentials"}
		logrus.Error(err.Error())
		return nil, err
	}

	jwt, err := utils.CreateJWT(id)

	if err != nil {
		return nil, err
	}

	return &models.UserAuth{
		Jwt: jwt,
	}, nil
}

// Register implements handlers.UserManager.
func (u user) Register(user models.User) (*models.UserAuth, error) {
	if u.repo.IsRegistered(user.Login) {
		err := servererrors.ConflictError{Used: user.Login}
		logrus.Error(err.Error())
		return nil, err
	}

	id, err := u.repo.Create(user)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	jwt, err := utils.CreateJWT(id)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return &models.UserAuth{
		Jwt: jwt,
	}, nil
}

func NewUser(repo UserRepository) user {
	return user{
		repo: repo,
	}
}
