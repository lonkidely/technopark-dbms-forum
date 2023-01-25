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
LIMIT 1
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
LIMIT 1
`
	CreateThread = `
INSERT INTO threads
(title, created, author, forum, message, slug)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`
)
