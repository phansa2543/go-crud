package usecases

import (
	"github.com/phansa2543/go-crud/entities"
	"github.com/phansa2543/go-crud/repositories"
)

type PostUsecase interface {
	PostUsecaseInsert(post *entities.Post) error
	PostUsecaseReadById(id string) (*entities.Post, error)
	PostUsecaseReadAllPost() ([]entities.Post, error)
	PostUsecaseUpdate(id string, data *entities.Post) error
	PostUsecaseDelete(id string) error
}

type postUsecaseImpl struct {
	postRepo repositories.PostRepository
}

func NewPostUsecaseImpl(postRepo repositories.PostRepository) PostUsecase {
	return &postUsecaseImpl{postRepo: postRepo}
}

func (p *postUsecaseImpl) PostUsecaseInsert(post *entities.Post) error {
	if err := p.postRepo.InsertPost(post); err != nil {
		return err
	}
	return nil
}

func (p *postUsecaseImpl) PostUsecaseReadById(id string) (*entities.Post, error) {
	post, err := p.postRepo.ReadById(id)
	if err != nil {
		return &entities.Post{}, err
	}
	return post, nil
}

func (p *postUsecaseImpl) PostUsecaseReadAllPost() ([]entities.Post, error) {
	post, err := p.postRepo.ReadAllPost()

	if err != nil {
		return post, err
	}
	return post, nil
}

func (p *postUsecaseImpl) PostUsecaseUpdate(id string, data *entities.Post) error {
	if err := p.postRepo.UpdatePost(id, data); err != nil {
		return err
	}
	return nil
}

func (p *postUsecaseImpl) PostUsecaseDelete(id string) error {
	if err := p.postRepo.DeletePost(id); err != nil {
		return err
	}
	return nil
}