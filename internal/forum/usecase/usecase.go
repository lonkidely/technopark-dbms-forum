package usecase

import (
	"database/sql"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"

	stdErrors "github.com/pkg/errors"

	"lonkidely/technopark-dbms-forum/internal/forum/repository"
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	userRepository "lonkidely/technopark-dbms-forum/internal/user/repository"
)

type ForumUsecase interface {
	CreateForum(forum *models.Forum) (models.Forum, error)
	GetForumDetails(forum *models.Forum) (models.Forum, error)
	GetForumThreads(forum *models.Forum, params *params.GetForumThreadsParams) ([]*models.Thread, error)
}

type forumUsecase struct {
	forumRepo repository.ForumRepository
	userRepo  userRepository.UserRepository
}

func NewForumUsecase(forumRepo repository.ForumRepository, userRepo userRepository.UserRepository) ForumUsecase {
	return &forumUsecase{
		forumRepo: forumRepo,
		userRepo:  userRepo,
	}
}

func (fu *forumUsecase) CreateForum(forum *models.Forum) (models.Forum, error) {
	forumExist, errorForumExist := fu.forumRepo.GetForumInfo(forum)
	if errorForumExist == nil {
		return forumExist, errors.ErrForumExist
	}

	userExist, errorUserExist := fu.userRepo.GetUserInfo(&models.User{Nickname: forum.User})
	if errorUserExist != nil {
		if stdErrors.Is(errorUserExist, sql.ErrNoRows) {
			return models.Forum{}, errors.ErrUserNotExist
		}
		return models.Forum{}, errorUserExist
	}

	forum.User = userExist.Nickname

	resultForum, err := fu.forumRepo.CreateForum(forum)
	return resultForum, err
}

func (fu *forumUsecase) GetForumDetails(forum *models.Forum) (models.Forum, error) {
	resultForum, err := fu.forumRepo.GetForumInfo(forum)
	if stdErrors.Is(err, sql.ErrNoRows) {
		return models.Forum{}, errors.ErrForumNotExist
	}
	if err != nil {
		return models.Forum{}, err
	}
	return resultForum, nil
}

func (fu *forumUsecase) GetForumThreads(forum *models.Forum, params *params.GetForumThreadsParams) ([]*models.Thread, error) {
	_, errExist := fu.forumRepo.GetForumInfo(forum)
	if stdErrors.Is(errExist, sql.ErrNoRows) {
		return []*models.Thread{}, errors.ErrForumNotExist
	}
	if errExist != nil {
		return []*models.Thread{}, errExist
	}

}
