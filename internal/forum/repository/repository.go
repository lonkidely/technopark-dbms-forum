package repository

import (
	"database/sql"
	"lonkidely/technopark-dbms-forum/internal/models"
)

type ForumRepository interface {
	GetForumInfo(forum *models.Forum) (models.Forum, error)
	CreateForum(forum *models.Forum) (models.Forum, error)
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
	response := models.Forum{}

	err := fr.db.QueryRow(GetForumInfo, forum.Slug).Scan(
		&response.Slug,
		&response.Title,
		&response.User,
		&response.Posts,
		&response.Threads)

	if err != nil {
		return models.Forum{}, err
	}

	return response, nil
}

func (fr *forumRepository) CreateForum(forum *models.Forum) (models.Forum, error) {
	response := models.Forum{}

	err := fr.db.QueryRow(CreateForum, forum.Slug, forum.Title, forum.User).Scan(
		&response.Slug,
		&response.Title,
		&response.User,
		&response.Posts,
		&response.Threads)

	if err != nil {
		return models.Forum{}, err
	}

	return response, nil
}
