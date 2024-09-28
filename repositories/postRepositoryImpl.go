package repositories

import (
	"github.com/phansa2543/go-crud/databases"
	"github.com/phansa2543/go-crud/entities"
)

type postPostgresRepository struct {
	db databases.Database
}

func NewPostPostgresRepository(db databases.Database) PostRepository {
	return &postPostgresRepository{db: db}
}

func (p *postPostgresRepository) InsertPost(post *entities.Post) error {

	if err := p.db.ConnectDB().Create(post); err != nil {
		return err.Error
	}

	return nil
}

func (p *postPostgresRepository) ReadById(id string) (*entities.Post, error) {

	var post entities.Post

	result := p.db.ConnectDB().Where("id = ?", id).First(&post)

	if result.Error != nil {
		return &entities.Post{}, result.Error
	}

	return &post, nil
}

func (p *postPostgresRepository) ReadAllPost() ([]entities.Post, error) {
	
	var post []entities.Post

	result := p.db.ConnectDB().Find(&post)

	if result.Error != nil {
		return post, result.Error
	}

	return post, nil
}

func (p *postPostgresRepository) UpdatePost(id string, data *entities.Post) error {
	var post entities.Post

	result := p.db.ConnectDB().Model(&post).Where("id = ?", id).Updates(data)
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *postPostgresRepository) DeletePost(id string) error {
	post := entities.Post{}

	result := p.db.ConnectDB().Where("id = ?", id).Delete(&post)

	if result.Error != nil {
		return result.Error
	}

	return nil
}