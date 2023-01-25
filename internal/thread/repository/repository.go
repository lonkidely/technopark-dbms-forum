package repository

import (
	"fmt"
	"strings"
	"time"

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
	GetThreadPostsFlat(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error)
	GetThreadPostsTree(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error)
	GetThreadPostsParent(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error)
	CreatePosts(thread *models.Thread, posts []*models.Post) ([]models.Post, error)
	GetParentPost(post *models.Post) (*models.Post, error)
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

func (tr *threadRepository) GetThreadPostsFlat(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error) {
	var query string
	var queryParams []interface{}

	if params.Since == -1 {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1
						ORDER BY id DESC
						LIMIT NULLIF($2, 0)`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1
						ORDER BY id
						LIMIT NULLIF($2, 0)`
		}
		queryParams = append(queryParams, thread.ID, params.Limit)
	} else {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1 AND id < $2
						ORDER BY id DESC
						LIMIT NULLIF($3, 0)`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1 AND id > $2
						ORDER BY id
						LIMIT NULLIF($3, 0)`
		}
		queryParams = append(queryParams, thread.ID, params.Since, params.Limit)
	}

	var response []models.Post

	row, err := tr.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var post models.Post
		errScan := row.Scan(
			&post.ID,
			&post.Author,
			&post.Created,
			&post.Forum,
			&post.Message,
			&post.Parent,
			&post.Thread,
			&post.IsEdited)

		if errScan != nil {
			return nil, errScan
		}

		response = append(response, post)
	}

	return response, nil
}

func (tr *threadRepository) GetThreadPostsTree(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error) {
	var query string
	var queryParams []interface{}

	if params.Since == -1 {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1
						ORDER BY path DESC, id DESC
						LIMIT NULLIF($2, 0)`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1
						ORDER BY path ASC, id ASC
						LIMIT NULLIF($2, 0)`
		}
		queryParams = append(queryParams, thread.ID, params.Limit)
	} else {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1 AND path < (SELECT path FROM posts WHERE id = $2)
						ORDER BY path DESC, id DESC
						LIMIT NULLIF($3, 0)`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE id_thread = $1 AND path > (SELECT path FROM posts WHERE id = $2)
						ORDER BY path ASC, id ASC
						LIMIT NULLIF($3, 0)`
		}
		queryParams = append(queryParams, thread.ID, params.Since, params.Limit)
	}

	var response []models.Post

	row, err := tr.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var post models.Post
		errScan := row.Scan(
			&post.ID,
			&post.Author,
			&post.Created,
			&post.Forum,
			&post.Message,
			&post.Parent,
			&post.Thread,
			&post.IsEdited)

		if errScan != nil {
			return nil, errScan
		}

		response = append(response, post)
	}

	return response, nil
}

func (tr *threadRepository) GetThreadPostsParent(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error) {
	var query string
	var queryParams []interface{}

	if params.Since == -1 {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE path[1] IN (SELECT id FROM posts WHERE id_thread = $1 AND parent = 0 ORDER BY id DESC LIMIT $2)
						ORDER BY path[1] DESC, path, id`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE path[1] IN (SELECT id FROM posts WHERE id_thread = $1 AND parent = 0 ORDER BY id LIMIT $2)
						ORDER BY path, id`
		}
		queryParams = append(queryParams, thread.ID, params.Limit)
	} else {
		if params.Desc {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE path[1] IN (SELECT id FROM posts WHERE id_thread = $1 AND parent = 0 AND PATH[1] <
						(SELECT path[1] FROM posts WHERE id = $2) ORDER BY id DESC LIMIT $3) 
						ORDER BY path[1] DESC, path, id`
		} else {
			query = `SELECT id, author, created, forum, message, parent, id_thread, isedited FROM posts
						WHERE path[1] IN (SELECT id FROM posts WHERE id_thread = $1 AND parent = 0 AND PATH[1] >
						(SELECT path[1] FROM posts WHERE id = $2) ORDER BY id ASC LIMIT $3) 
						ORDER BY path, id`
		}
		queryParams = append(queryParams, thread.ID, params.Since, params.Limit)
	}

	var response []models.Post

	row, err := tr.db.Query(query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	for row.Next() {
		var post models.Post
		errScan := row.Scan(
			&post.ID,
			&post.Author,
			&post.Created,
			&post.Forum,
			&post.Message,
			&post.Parent,
			&post.Thread,
			&post.IsEdited)

		if errScan != nil {
			return nil, errScan
		}

		response = append(response, post)
	}

	return response, nil
}

func (tr *threadRepository) GetParentPost(post *models.Post) (*models.Post, error) {
	query := `SELECT id_thread FROM posts WHERE id = $1`

	response := &models.Post{}

	err := tr.db.QueryRow(query, post.Parent).Scan(&response.Thread)

	if stdErrors.Is(err, pgx.ErrNoRows) {
		return nil, errors.ErrThreadExist
	}

	return response, nil
}

func (tr *threadRepository) CreatePosts(thread *models.Thread, posts []*models.Post) ([]models.Post, error) {
	query := `INSERT INTO posts(author, created, forum, message, parent, id_thread) VALUES `

	var queryParameters []interface{}
	created := time.Now()
	for i, p := range posts {
		value := fmt.Sprintf(
			"($%d, $%d, $%d, $%d, $%d, $%d),",
			i*6+1, i*6+2, i*6+3, i*6+4, i*6+5, i*6+6,
		)

		query += value
		queryParameters = append(queryParameters, p.Author, created, thread.Forum, p.Message, p.Parent, thread.ID)
	}

	query = strings.TrimSuffix(query, ",")
	query += ` RETURNING id, author, created, forum, message, parent, id_thread, isedited`

	row, err := tr.db.Query(query, queryParameters...)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var response []models.Post

	for row.Next() {
		post := models.Post{}

		errScan := row.Scan(
			&post.ID,
			&post.Author,
			&post.Created,
			&post.Forum,
			&post.Message,
			&post.Parent,
			&post.Thread,
			&post.IsEdited)

		if errScan != nil {
			return nil, errScan
		}

		response = append(response, post)
	}

	return response, err
}
