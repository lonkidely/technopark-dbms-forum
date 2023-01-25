package params

type GetForumThreadsParams struct {
	Limit int
	Since string
	Desc  bool
}

type GetForumUsersParams struct {
	Limit int
	Since string
	Desc  bool
}

type VoteThreadParams struct {
	Nickname string
	Voice    int
}

type GetPostsParams struct {
	Limit int
	Since int
	Desc  bool
	Sort  string
}
