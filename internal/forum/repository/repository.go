package repository

import (
	"database/sql"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

type ForumRepository interface {
	GetForumInfo(forum *models.Forum) (models.Forum, error)
	CreateForum(forum *models.Forum) (models.Forum, error)
	GetForumThreads(forum *models.Forum, params *params.GetForumThreadsParams) ([]*models.Thread, error)
	GetForumUsers(forum *models.Forum, params *params.GetForumUsersParams) ([]*models.User, error)
}

type forumRepository struct {
	db *sql.DB
}

func NewForumRepository(db *sql.DB) ForumRepository {
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

	if err != nil {
		return models.Forum{}, err
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
		currentThread := &models.Thread{}
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

		result = append(result, currentThread)
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
