package repository

import (
	"github.com/jackc/pgx"

	"lonkidely/technopark-dbms-forum/internal/models"
)

type ServiceRepository interface {
	Clear() error
	Status() (models.Status, error)
}

type serviceRepository struct {
	db *pgx.ConnPool
}

func NewServiceRepository(db *pgx.ConnPool) ServiceRepository {
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
	err := row.Scan(&response.Forum)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountPosts)
	err = row.Scan(&response.Post)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountThreads)
	err = row.Scan(&response.Thread)
	if err != nil {
		return models.Status{}, err
	}

	row = sr.db.QueryRow(CountUsers)
	err = row.Scan(&response.User)
	if err != nil {
		return models.Status{}, err
	}

	return response, nil
}
