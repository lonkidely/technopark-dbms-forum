package repository

const (
	GetForumInfo = `
SELECT
    slug, title, "user", posts, threads
FROM forum
WHERE slug = $1
`

	CreateForum = `
INSERT INTO forum
(slug, title, "user")
VALUES ($1, $2, $3)
RETURNING slug, title, "user", posts, threads
`

	GetForumThreadsAsc = `
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
WHERE forum = $1
ORDER BY created
LIMIT $2
`
	GetForumThreadsDesc = `
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
WHERE forum = $1
ORDER BY created DESC
LIMIT $2
`
	GetForumThreadsSinceAsc = `
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
WHERE forum = $1 AND created >= $2
ORDER BY created
LIMIT $3
`
	GetForumThreadsSinceDesc = `
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
WHERE forum = $1 AND created <= $2
ORDER BY created DESC
LIMIT $3
`

	GetForumUsersAsc = `
SELECT
    nickname,
    fullname,
    about,
    email
FROM forum_users
WHERE forum = $1
ORDER BY nickname
LIMIT $2
`
	GetForumUsersDesc = `
SELECT
    nickname,
    fullname,
    about,
    email
FROM forum_users
WHERE forum = $1
ORDER BY nickname DESC
LIMIT $2
`
	GetForumUsersSinceAsc = `
SELECT
    nickname,
    fullname,
    about,
    email
FROM forum_users
WHERE forum = $1 AND nickname > $2
ORDER BY nickname
LIMIT $3
`
	GetForumUsersSinceDesc = `
SELECT
    nickname,
    fullname,
    about,
    email
FROM forum_users
WHERE forum = $1 AND nickname < $2
ORDER BY nickname DESC
LIMIT $3
`
)
