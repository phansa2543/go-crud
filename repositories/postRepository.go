package repositories

import (
	"github.com/phansa2543/go-crud/entities"
)

type PostRepository interface {
	InsertPost(post *entities.Post) error
	ReadById(id string) (*entities.Post, error)
	ReadAllPost() ([]entities.Post, error)
	UpdatePost(id string, data *entities.Post) error
	DeletePost(id string) error
}