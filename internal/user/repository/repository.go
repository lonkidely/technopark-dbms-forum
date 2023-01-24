package repository

import (
	"database/sql"

	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
)

type UserRepository interface {
	GetUsers(user *models.User) ([]models.User, error)
	CreateUser(user *models.User) (models.User, error)
	GetUserInfo(user *models.User) (models.User, error)
	CheckUserExist(user *models.User) (bool, error)
	CheckEmailIsNotUsed(user *models.User) (bool, error)
	UpdateUser(user *models.User) (models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) GetUsers(user *models.User) ([]models.User, error) {
	result := make([]models.User, 0)

	rows, err := ur.db.Query(GetUser, user.Nickname, user.Email)
	if err != nil {
		return result, errors.ErrUserNotExist
	}
	defer rows.Close()

	for rows.Next() {
		currentUser := models.User{}

		err = rows.Scan(
			&currentUser.Nickname,
			&currentUser.FullName,
			&currentUser.Email,
			&currentUser.About)

		if err != nil {
			return result, err
		}

		result = append(result, currentUser)
	}

	if len(result) == 0 {
		return []models.User{}, errors.ErrUserNotExist
	}

	return result, nil
}

func (ur *userRepository) CreateUser(user *models.User) (models.User, error) {
	row := ur.db.QueryRow(CreateUser, user.Nickname, user.FullName, user.Email, user.About)
	if row.Err() != nil {
		return models.User{}, row.Err()
	}
	return *user, nil
}

func (ur *userRepository) GetUserInfo(user *models.User) (models.User, error) {
	response := models.User{}

	err := ur.db.QueryRow(GetUserInfo, user.Nickname).Scan(&response.Nickname,
		&response.FullName,
		&response.Email,
		&response.About)

	if stdErrors.Is(err, sql.ErrNoRows) {
		return models.User{}, errors.ErrUserNotExist
	}
	if err != nil {
		return models.User{}, err
	}

	return response, nil
}

func (ur *userRepository) CheckUserExist(user *models.User) (bool, error) {
	response := false
	err := ur.db.QueryRow(CheckUserExist, user.Nickname).Scan(&response)
	return response, err
}

func (ur *userRepository) CheckEmailIsNotUsed(user *models.User) (bool, error) {
	response := true
	err := ur.db.QueryRow(CheckEmailNotUsed, user.Email, user.Nickname).Scan(&response)
	return response, err
}

func (ur *userRepository) UpdateUser(user *models.User) (models.User, error) {
	response := models.User{}

	err := ur.db.QueryRow(UpdateUser, user.Nickname, user.FullName, user.Email, user.About).Scan(
		&response.Nickname,
		&response.FullName,
		&response.Email,
		&response.About)

	if err != nil {
		return models.User{}, err
	}

	return response, nil
}
