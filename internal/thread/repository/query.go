package repository

const (
	CheckThreadExist    = `SELECT EXISTS (SELECT 1 FROM threads WHERE slug = $1)`
	GetThreadInfoBySlug = `
SELECT 
    id,
    title,
    created,
    author,
    forum,
    message,
    slug,
    votes
FROM threads
WHERE slug = $1
`
	GetThreadInfoByID = `
SELECT
    id,
    title,
    created,
    author,
    forum,
    message,
    slug,
    votes
FROM threads
WHERE id = $1
`
	CreateThread = `
INSERT INTO threads
(title, created, author, forum, message, slug)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`
	UpdateThreadByID = `
UPDATE threads
SET title = COALESCE(NULLIF($2, ''), title), message = COALESCE(NULLIF($3, ''), message)
WHERE id = $1
RETURNING id, title, created, author, forum, message, slug, votes
`
	UpdateThreadBySlug = `
UPDATE threads
SET title = COALESCE(NULLIF($2, ''), title), message = COALESCE(NULLIF($3, ''), message)
WHERE slug = $1
RETURNING id, title, created, author, forum, message, slug, votes
`
	CheckExistVote = `
SELECT EXISTS(SELECT 1 FROM votes WHERE author = $1 AND id_thread = $2)
`
	InsertVote = `
INSERT INTO votes
(author, voice, id_thread)
VALUES ($1, $2, $3)
`
	UpdateVote = `
UPDATE votes
SET voice = $1
WHERE author = $2 AND id_thread = $3
`
)
