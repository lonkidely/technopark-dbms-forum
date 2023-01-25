package repository

import (
	"github.com/jackc/pgx"

	"lonkidely/technopark-dbms-forum/internal/models"
)

type ThreadRepository interface {
	CreateThread(thread *models.Thread) (models.Thread, error)
	GetThreadInfo(thread *models.Thread) (models.Thread, error)
}

type threadRepository struct {
	db *pgx.ConnPool
}

func NewThreadRepository(db *pgx.ConnPool) ThreadRepository {
	return &threadRepository{
		db: db,
	}
}

func (tr *threadRepository) CreateThread(thread *models.Thread) (models.Thread, error) {
	row := tr.db.QueryRow(CreateThread, thread.Title, thread.Created, thread.Author, thread.Forum, thread.Message, thread.Slug)
	err := row.Scan(&thread.ID)
	if err != nil {
		return models.Thread{}, err
	}

	return *thread, nil
}

func (tr *threadRepository) GetThreadInfo(thread *models.Thread) (models.Thread, error) {
	response := models.Thread{}

	row := tr.db.QueryRow(GetThreadInfoBySlug, thread.Slug)

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
