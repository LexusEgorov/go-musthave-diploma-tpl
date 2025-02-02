package user

import (
	"time"

	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/db"
	"github.com/LexusEgorov/go-musthave-diploma-tpl/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"
)

type userRepo struct {
	db   db.DB
	psql squirrel.StatementBuilderType
}

// Create implements UserRepository.
func (u userRepo) Create(user models.User) (int, error) {
	currTime := time.Now().String()
	sql, args, err := u.psql.Insert("users").
		Columns("login", "password", "created_at", "updated_at").
		Values(user.Login).Values(user.Password).Values(currTime).Values(currTime).ToSql()

	if err != nil {
		return 0, err
	}

	res, err := u.db.DB.Exec(sql, args...)

	if err != nil {
		return 0, err
	}

	uId, err := res.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(uId), nil
}

// FindById implements UserRepository.
func (u userRepo) FindById(id uint) (*models.User, error) {
	user := models.User{}
	sql, args, err := u.psql.Select("login", "password", "balance").From("users").Where("uId = ?", id).ToSql()

	if err != nil {
		return nil, err
	}

	err = u.db.DB.QueryRow(sql, args...).Scan(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// IsRegistered implements UserRepository.
func (u userRepo) IsRegistered(login string) bool {
	user := models.User{}
	sql, args, err := u.psql.Select("*").From("users").Where("login = ?", login).ToSql()

	if err != nil {
		logrus.Error(err)
		return false
	}

	err = u.db.DB.QueryRow(sql, args...).Scan(&user)

	if err != nil {
		logrus.Error(err)
		return false
	}

	return user.Login != ""
}

func NewUserRepo(db db.DB) UserRepository {
	return userRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
