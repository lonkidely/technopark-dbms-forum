package usecase

import (
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/params"
	"lonkidely/technopark-dbms-forum/internal/post/repository"
)

type PostUsecase interface {
	GetPostDetails(post *models.Post, params *params.PostDetailsParams) (*models.PostDetails, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}

type postUsecase struct {
	postRepo repository.PostRepository
}

func NewPostUsecase(pr repository.PostRepository) PostUsecase {
	return &postUsecase{
		postRepo: pr,
	}
}

func (pu *postUsecase) GetPostDetails(post *models.Post, params *params.PostDetailsParams) (*models.PostDetails, error) {
	return pu.postRepo.GetPostDetails(post, params)
}

func (pu *postUsecase) UpdatePost(post *models.Post) (*models.Post, error) {
	if post.Message == "" {
		resPost, err := pu.GetPostDetails(post, &params.PostDetailsParams{})
		if err != nil {
			return nil, err
		}
		return &resPost.Post, nil
	}
	return pu.postRepo.UpdatePost(post)
}
