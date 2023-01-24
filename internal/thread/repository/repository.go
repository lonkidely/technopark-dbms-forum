package repository

import (
	"database/sql"
	"time"

	"lonkidely/technopark-dbms-forum/internal/models"
)

type ThreadRepository interface {
	CreateThread(thread *models.Thread) (models.Thread, error)
	GetThreadInfo(thread *models.Thread) (models.Thread, error)
}

type threadRepository struct {
	db *sql.DB
}

func NewThreadRepository(db *sql.DB) ThreadRepository {
	return &threadRepository{
		db: db,
	}
}

func (tr *threadRepository) CreateThread(thread *models.Thread) (models.Thread, error) {
	if thread.Created == "" {
		thread.Created = time.Now().Format(time.RFC3339)
	}

	row := tr.db.QueryRow(CreateThread, thread.Title, thread.Created, thread.Author, thread.Forum, thread.Message, thread.Slug)
	if row.Err() != nil {
		return models.Thread{}, row.Err()
	}
	err := row.Scan(&thread.ID)
	if err != nil {
		return models.Thread{}, err
	}

	return *thread, nil
}

func (tr *threadRepository) GetThreadInfo(thread *models.Thread) (models.Thread, error) {
	response := models.Thread{}

	row := tr.db.QueryRow(GetThreadInfoBySlug, thread.Slug)
	if row.Err() != nil {
		return models.Thread{}, row.Err()
	}

	err := row.Scan(
		&response.ID,
		&response.Title,
		&response.Created,
		&response.Author,
		&response.Forum,
		&response.Message,
		&response.Slug,
		&response.Votes)
	if err != nil {
		return models.Thread{}, err
	}

	return response, nil
}
