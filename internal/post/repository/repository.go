package repository

import (
	"github.com/jackc/pgx"
	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
)

type PostRepository interface {
	GetPostDetails(post *models.Post, params *params.PostDetailsParams) (*models.PostDetails, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}

type postRepository struct {
	db *pgx.ConnPool
}

func NewPostRepository(db *pgx.ConnPool) PostRepository {
	return &postRepository{
		db: db,
	}
}

func (pr *postRepository) GetPostDetails(post *models.Post, params *params.PostDetailsParams) (*models.PostDetails, error) {
	response := &models.PostDetails{}
	response.Post.ID = post.ID

	err := pr.db.QueryRow(GetPost, post.ID).Scan(
		&response.Post.Author,
		&response.Post.Created,
		&response.Post.Forum,
		&response.Post.Message,
		&response.Post.Parent,
		&response.Post.Thread,
		&response.Post.IsEdited)

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return nil, errors.ErrPostNotExist
	}
	if err != nil {
		return nil, err
	}

	for _, val := range params.Related {
		switch val {
		case "forum":
			err = pr.db.QueryRow(GetPostForum, response.Post.Forum).Scan(
				&response.Forum.Slug,
				&response.Forum.Title,
				&response.Forum.User,
				&response.Forum.Posts,
				&response.Forum.Threads)
			if stdErrors.Is(err, pgx.ErrNoRows) {
				return nil, errors.ErrPostNotExist
			}
			if err != nil {
				return nil, err
			}
		case "thread":
			err = pr.db.QueryRow(GetPostThread, response.Post.Thread).Scan(
				&response.Thread.ID,
				&response.Thread.Title,
				&response.Thread.Created,
				&response.Thread.Author,
				&response.Thread.Forum,
				&response.Thread.Message,
				&response.Thread.Slug,
				&response.Thread.Votes)
			if stdErrors.Is(err, pgx.ErrNoRows) {
				return nil, errors.ErrPostNotExist
			}
			if err != nil {
				return nil, err
			}
		case "user":
			err = pr.db.QueryRow(GetPostUser, response.Post.Author).Scan(
				&response.Author.Nickname,
				&response.Author.FullName,
				&response.Author.Email,
				&response.Author.About)
			if stdErrors.Is(err, pgx.ErrNoRows) {
				return nil, errors.ErrPostNotExist
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return response, nil
}

func (pr *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	response := &models.Post{}

	err := pr.db.QueryRow(UpdatePost, post.ID, post.Message).Scan(
		&response.Author,
		&response.Created,
		&response.Forum,
		&response.Message,
		&response.Parent,
		&response.Thread,
		&response.IsEdited)

	response.ID = post.ID

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return nil, errors.ErrPostNotExist
	}
	if err != nil {
		return nil, err
	}

	return response, nil
}
