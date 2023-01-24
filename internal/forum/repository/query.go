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
)
