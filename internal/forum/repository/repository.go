package repository

import (
	"github.com/jackc/pgx"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

type ForumRepository interface {
	GetForumInfo(forum *models.Forum) (models.Forum, error)
	CheckForumExist(forum *models.Forum) (bool, error)
	CreateForum(forum *models.Forum) (models.Forum, error)
	GetForumThreads(forum *models.Forum, params *params.GetForumThreadsParams) ([]*models.Thread, error)
	GetForumUsers(forum *models.Forum, params *params.GetForumUsersParams) ([]*models.User, error)
}

type forumRepository struct {
	db *pgx.ConnPool
}

func NewForumRepository(db *pgx.ConnPool) ForumRepository {
	return &forumRepository{
		db: db,
	}
}

func (fr *forumRepository) GetForumInfo(forum *models.Forum) (models.Forum, error) {
	result := models.Forum{}

	err := fr.db.QueryRow(GetForumInfo, forum.Slug).Scan(
		&result.Slug,
		&result.Title,
		&result.User,
		&result.Posts,
		&result.Threads)

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return models.Forum{}, errors.ErrForumNotExist
	}

	if err != nil {
		return models.Forum{}, err
	}

	return result, nil
}

func (fr *forumRepository) CheckForumExist(forum *models.Forum) (bool, error) {
	result := false

	row := fr.db.QueryRow(CheckForumExist)

	err := row.Scan(&result)
	if err != nil {
		return false, err
	}

	return result, nil
}

func (fr *forumRepository) CreateForum(forum *models.Forum) (models.Forum, error) {
	result := models.Forum{}

	err := fr.db.QueryRow(CreateForum, forum.Slug, forum.Title, forum.User).Scan(
		&result.Slug,
		&result.Title,
		&result.User,
		&result.Posts,
		&result.Threads)

	if err != nil {
		return models.Forum{}, err
	}

	return result, nil
}

func (fr *forumRepository) GetForumThreads(forum *models.Forum, params *params.GetForumThreadsParams) ([]*models.Thread, error) {
	var query string
	var queryParams []interface{}

	if params.Since == "" {
		if params.Desc {
			query = GetForumThreadsDesc
		} else {
			query = GetForumThreadsAsc
		}
		queryParams = append(queryParams, forum.Slug, params.Limit)
	} else {
		if params.Desc {
			query = GetForumThreadsSinceDesc
		} else {
			query = GetForumThreadsSinceAsc
		}
		queryParams = append(queryParams, forum.Slug, params.Since, params.Limit)
	}

	rows, err := fr.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.Thread

	for rows.Next() {
		currentThread := models.Thread{}
		errScan := rows.Scan(
			&currentThread.ID,
			&currentThread.Title,
			&currentThread.Created,
			&currentThread.Author,
			&currentThread.Forum,
			&currentThread.Message,
			&currentThread.Slug,
			&currentThread.Votes)

		if errScan != nil {
			return nil, err
		}

		result = append(result, &currentThread)
	}

	return result, nil
}

func (fr *forumRepository) GetForumUsers(forum *models.Forum, params *params.GetForumUsersParams) ([]*models.User, error) {
	var query string
	var queryParams []interface{}

	if params.Since == "" {
		if params.Desc {
			query = GetForumUsersDesc
		} else {
			query = GetForumUsersAsc
		}
		queryParams = append(queryParams, forum.Slug, params.Limit)
	} else {
		if params.Desc {
			query = GetForumUsersSinceDesc
		} else {
			query = GetForumUsersSinceAsc
		}
		queryParams = append(queryParams, forum.Slug, params.Since, params.Limit)
	}

	rows, err := fr.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []*models.User

	for rows.Next() {
		currentUser := &models.User{}
		errScan := rows.Scan(
			&currentUser.Nickname,
			&currentUser.FullName,
			&currentUser.About,
			&currentUser.Email)

		if errScan != nil {
			return nil, err
		}

		result = append(result, currentUser)
	}

	return result, nil
}
