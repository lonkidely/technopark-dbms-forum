package repository

const (
	GetUser = `
SELECT
    nickname,
    fullname,
    email,
    about
FROM users
WHERE nickname = $1 OR email = $2
`

	CreateUser = `
INSERT INTO users
(nickname, fullname, email, about)
VALUES ($1, $2, $3, $4)
`

	GetUserInfo = `
SELECT
    nickname,
    fullname,
    email,
    about
FROM users
WHERE nickname = $1
`

	CheckUserExist = `SELECT EXISTS (SELECT 1 FROM users WHERE nickname = $1)`

	CheckEmailNotUsed = `SELECT EXISTS (SELECT 1 FROM users WHERE email = $1 AND nickname != $2)`

	UpdateUser = `
UPDATE users
SET 
    fullname = COALESCE(NULLIF($2, ''), fullname), 
    email = COALESCE(NULLIF($3, ''), email), 
    about = COALESCE(NULLIF($4, ''), about)
WHERE nickname = $1
RETURNING nickname, fullname, email, about
`
)
