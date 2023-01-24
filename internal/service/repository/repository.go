package repository

import (
	"database/sql"

	"lonkidely/technopark-dbms-forum/internal/models"
)

type ServiceRepository interface {
	Clear() error
	Status() (models.Status, error)
}

type serviceRepository struct {
	db *sql.DB
}

func NewServiceRepository(db *sql.DB) ServiceRepository {
	return &serviceRepository{
		db: db,
	}
}

func (sr *serviceRepository) Clear() error {
	_, err := sr.db.Exec(Clear)
	return err
}

func (sr *serviceRepository) Status() (models.Status, error) {
	response := models.Status{}

	row := sr.db.QueryRow(CountForums)
	if row.Err() != nil {
		return models.Status{}, row.Err()
	}
	err := row.Scan(&response.Forum)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountPosts)
	if row.Err() != nil {
		return models.Status{}, row.Err()
	}
	err = row.Scan(&response.Post)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountThreads)
	if row.Err() != nil {
		return models.Status{}, row.Err()
	}
	err = row.Scan(&response.Thread)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountUsers)
	if row.Err() != nil {
		return models.Status{}, row.Err()
	}
	err = row.Scan(&response.User)
	if err != nil {
		return models.Status{}, err
	}

	return response, nil
}
