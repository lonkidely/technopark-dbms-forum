package repository

import (
	"github.com/jackc/pgx"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

type ThreadRepository interface {
	CreateThread(thread *models.Thread) (models.Thread, error)
	GetThreadBySlug(thread *models.Thread) (models.Thread, error)
	GetThreadByID(thread *models.Thread) (models.Thread, error)
	UpdateThreadDetails(thread *models.Thread) (models.Thread, error)
	CheckExistVote(thread *models.Thread, params *params.VoteThreadParams) (bool, error)
	InsertVoteThread(thread *models.Thread, params *params.VoteThreadParams) error
	UpdateVoteThread(thread *models.Thread, params *params.VoteThreadParams) error
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

func (tr *threadRepository) GetThreadBySlug(thread *models.Thread) (models.Thread, error) {
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

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return models.Thread{}, errors.ErrThreadNotExist
	}

	if err != nil {
		return models.Thread{}, err
	}

	return response, nil
}

func (tr *threadRepository) GetThreadByID(thread *models.Thread) (models.Thread, error) {
	response := models.Thread{}

	row := tr.db.QueryRow(GetThreadInfoByID, thread.ID)

	err := row.Scan(
		&response.ID,
		&response.Title,
		&response.Created,
		&response.Author,
		&response.Forum,
		&response.Message,
		&response.Slug,
		&response.Votes)

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return models.Thread{}, errors.ErrThreadNotExist
	}

	if err != nil {
		return models.Thread{}, err
	}

	return response, nil
}

func (tr *threadRepository) UpdateThreadDetails(thread *models.Thread) (models.Thread, error) {
	var query string
	var queryParams []interface{}
	if thread.Slug != "" {
		query = UpdateThreadBySlug
		queryParams = append(queryParams, thread.Slug)
	} else {
		query = UpdateThreadByID
		queryParams = append(queryParams, thread.ID)
	}
	queryParams = append(queryParams, thread.Title, thread.Message)

	response := models.Thread{}

	err := tr.db.QueryRow(query, queryParams...).Scan(
		&response.ID,
		&response.Title,
		&response.Created,
		&response.Author,
		&response.Forum,
		&response.Message,
		&response.Slug,
		&response.Votes)

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return models.Thread{}, errors.ErrThreadNotExist
	}

	if err != nil {
		return models.Thread{}, err
	}

	return response, nil
}

func (tr *threadRepository) CheckExistVote(thread *models.Thread, params *params.VoteThreadParams) (bool, error) {
	response := true

	err := tr.db.QueryRow(CheckExistVote, params.Nickname, thread.ID).Scan(&response)

	return response, err
}

func (tr *threadRepository) InsertVoteThread(thread *models.Thread, params *params.VoteThreadParams) error {
	_, err := tr.db.Exec(InsertVote, params.Nickname, params.Voice, thread.ID)
	return err
}

func (tr *threadRepository) UpdateVoteThread(thread *models.Thread, params *params.VoteThreadParams) error {
	_, err := tr.db.Exec(UpdateVote, params.Voice, params.Nickname, thread.ID)
	return err
}
