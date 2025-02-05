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

// Auth implements UserRepository.
func (u userRepo) Auth(user models.User) (int, bool) {
	sql, args, err := u.psql.Select("id").
		From("users").
		Where("login = ?", user.Login).
		Where("password = ?", user.Password).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return 0, false
	}

	rows, err := u.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return 0, false
	}

	if rows.Next() {
		var uID int
		rows.Scan(&uID)

		return uID, true
	}

	return 0, false
}

// Create implements UserRepository.
func (u userRepo) Create(user models.User) (int, error) {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	sql, args, err := u.psql.Insert("users").
		Columns("login", "password", "created_at", "updated_at").
		Values(user.Login, user.Password, currTime, currTime).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	var uID int
	err = u.db.DB.QueryRow(sql, args...).Scan(&uID)

	if err != nil {
		logrus.Error(err)
		return 0, err
	}

	return int(uID), nil
}

// FindById implements UserRepository.
func (u userRepo) FindByID(id uint) (*models.User, error) {
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
	sql, args, err := u.psql.Select("*").
		From("users").
		Where("login = ?", login).
		ToSql()

	if err != nil {
		logrus.Error(err)
		return false
	}

	rows, err := u.db.DB.Query(sql, args...)

	if err != nil {
		logrus.Error(err)
		return false
	}

	return rows.Next()
}

func NewUserRepo(db db.DB) UserRepository {
	return userRepo{
		db:   db,
		psql: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
