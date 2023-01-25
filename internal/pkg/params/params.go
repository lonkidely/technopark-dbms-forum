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
