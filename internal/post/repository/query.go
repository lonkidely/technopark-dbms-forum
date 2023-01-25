package repository

const (
	GetPost = `
SELECT author,
       created,
       forum,
       message,
       parent,
       id_thread,
       isedited
from posts
WHERE id = $1
`
	GetPostForum = `
SELECT
    slug, title, "user", posts, threads
FROM forum
WHERE slug = $1
`
	GetPostThread = `
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
	GetPostUser = `
SELECT
    nickname,
    fullname,
    email,
    about
FROM users
WHERE nickname = $1
`
	UpdatePost = `
UPDATE posts
SET message  = COALESCE(NULLIF($2, ''), message),
    isedited = CASE
				   WHEN $2 = message THEN isedited
				   ELSE true
    END
WHERE id = $1
RETURNING author, created, forum, message, parent, id_thread, isedited
`
)
