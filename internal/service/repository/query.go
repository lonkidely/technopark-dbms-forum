package repository

const (
	Clear        = `TRUNCATE threads, users, forum, forum_users, posts, votes CASCADE`
	CountUsers   = `SELECT count(*) FROM users`
	CountForums  = `SELECT count(*) FROM forum`
	CountThreads = `SELECT count(*) FROM threads`
	CountPosts   = `SELECT count(*) FROM posts`
)
