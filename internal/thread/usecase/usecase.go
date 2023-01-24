package usecase

import (
	forumRepository "lonkidely/technopark-dbms-forum/internal/forum/repository"
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/thread/repository"
	userRepository "lonkidely/technopark-dbms-forum/internal/user/repository"
)

type ThreadUsecase interface {
	CreateThread(thread *models.Thread) (models.Thread, error)
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
		threadExist, errThreadExist := tu.threadRepo.GetThreadInfo(thread)
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
