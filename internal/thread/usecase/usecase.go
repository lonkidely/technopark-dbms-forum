package usecase

import (
	forumRepository "lonkidely/technopark-dbms-forum/internal/forum/repository"
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
	"lonkidely/technopark-dbms-forum/internal/thread/repository"
	userRepository "lonkidely/technopark-dbms-forum/internal/user/repository"
)

type ThreadUsecase interface {
	CreateThread(thread *models.Thread) (models.Thread, error)
	GetThreadDetails(thread *models.Thread) (models.Thread, error)
	UpdateThreadDetails(thread *models.Thread) (models.Thread, error)
	VoteThread(thread *models.Thread, params *params.VoteThreadParams) (models.Thread, error)
	GetPosts(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error)
	CreatePosts(thread *models.Thread, posts []*models.Post) ([]models.Post, error)
}

type threadUsecase struct {
	threadRepo repository.ThreadRepository
	forumRepo  forumRepository.ForumRepository
	userRepo   userRepository.UserRepository
}

func NewThreadUsecase(threadRepo repository.ThreadRepository, forumRepo forumRepository.ForumRepository, userRepo userRepository.UserRepository) ThreadUsecase {
	return &threadUsecase{
		threadRepo: threadRepo,
		forumRepo:  forumRepo,
		userRepo:   userRepo,
	}
}

func (tu *threadUsecase) CreateThread(thread *models.Thread) (models.Thread, error) {
	forumExist, errForumExist := tu.forumRepo.GetForumInfo(&models.Forum{Slug: thread.Forum})
	if errForumExist != nil {
		return models.Thread{}, errForumExist
	}
	thread.Forum = forumExist.Slug

	userExist, errUserExist := tu.userRepo.GetUserInfo(&models.User{Nickname: thread.Author})
	if errUserExist != nil {
		return models.Thread{}, errUserExist
	}
	thread.Author = userExist.Nickname

	if thread.Slug != "" {
		threadExist, errThreadExist := tu.threadRepo.GetThreadBySlug(thread)
		if errThreadExist == nil {
			return threadExist, errors.ErrThreadExist
		}
	}

	resultThread, err := tu.threadRepo.CreateThread(thread)
	if err != nil {
		return models.Thread{}, err
	}
	return resultThread, nil
}

func (tu *threadUsecase) GetThreadDetails(thread *models.Thread) (models.Thread, error) {
	if thread.Slug != "" {
		return tu.threadRepo.GetThreadBySlug(thread)
	}
	return tu.threadRepo.GetThreadByID(thread)
}

func (tu *threadUsecase) UpdateThreadDetails(thread *models.Thread) (models.Thread, error) {
	return tu.threadRepo.UpdateThreadDetails(thread)
}

func (tu *threadUsecase) VoteThread(thread *models.Thread, params *params.VoteThreadParams) (models.Thread, error) {
	currentThread, errThread := tu.GetThreadDetails(thread)
	if errThread != nil {
		return models.Thread{}, errThread
	}

	currentUser, errUser := tu.userRepo.GetUserInfo(&models.User{Nickname: params.Nickname})
	if errUser != nil {
		return models.Thread{}, errUser
	}
	params.Nickname = currentUser.Nickname

	voteExist, errVoteExist := tu.threadRepo.CheckExistVote(&currentThread, params)
	if errVoteExist != nil {
		return models.Thread{}, errVoteExist
	}

	if voteExist {
		err := tu.threadRepo.UpdateVoteThread(&currentThread, params)
		if err != nil {
			return models.Thread{}, err
		}
	} else {
		err := tu.threadRepo.InsertVoteThread(&currentThread, params)
		if err != nil {
			return models.Thread{}, err
		}
	}

	return tu.GetThreadDetails(&currentThread)
}

func (tu *threadUsecase) GetPosts(thread *models.Thread, params *params.GetPostsParams) ([]models.Post, error) {
	currentThread, threadErr := tu.GetThreadDetails(thread)
	if threadErr != nil {
		return nil, threadErr
	}

	if params.Sort == "flat" {
		return tu.threadRepo.GetThreadPostsFlat(&currentThread, params)
	}
	if params.Sort == "tree" {
		return tu.threadRepo.GetThreadPostsTree(&currentThread, params)
	}
	return tu.threadRepo.GetThreadPostsParent(&currentThread, params)
}

func (tu *threadUsecase) CreatePosts(thread *models.Thread, posts []*models.Post) ([]models.Post, error) {
	currentThread, threadErr := tu.GetThreadDetails(thread)
	if threadErr != nil {
		return nil, threadErr
	}

	if len(posts) == 0 {
		return []models.Post{}, nil
	}

	currentUser, errUser := tu.userRepo.GetUserInfo(&models.User{Nickname: posts[0].Author})
	if errUser != nil {
		return nil, errUser
	}

	posts[0].Author = currentUser.Nickname

	if posts[0].Parent != 0 {
		var postWithParent *models.Post
		var postErr error

		postWithParent, postErr = tu.threadRepo.GetParentPost(posts[0])
		if postErr != nil {
			return []models.Post{}, postErr
		}

		if postWithParent.Thread != currentThread.ID {
			return nil, errors.ErrThreadExist
		}
	}

	resultPosts, err := tu.threadRepo.CreatePosts(&currentThread, posts)
	if err != nil {
		return []models.Post{}, err
	}

	return resultPosts, nil
}
