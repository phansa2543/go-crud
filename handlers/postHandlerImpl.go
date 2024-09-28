package handlers

import (

	"github.com/gofiber/fiber/v2"
	"github.com/phansa2543/go-crud/entities"
	"github.com/phansa2543/go-crud/usecases"
)

type postHandler struct {
	postUsecase usecases.PostUsecase
}

func NewPostHandler(postUsecase usecases.PostUsecase) PostHandler {
	return &postHandler{postUsecase: postUsecase}
}

func (p *postHandler) CreatePost(c *fiber.Ctx) error {
	post := new(entities.Post)

	if err := c.BodyParser(post); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := p.postUsecase.PostUsecaseInsert(post); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (p *postHandler) ReadById(c *fiber.Ctx) error {

	queryParams := c.Query("id")

	if queryParams == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	post, err := p.postUsecase.PostUsecaseReadById(queryParams)
	
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(post)
}

func (p *postHandler) ReadAllTitle(c *fiber.Ctx) error {
	post, err := p.postUsecase.PostUsecaseReadAllPost()

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(post)
}

func (p *postHandler) UpdatePost(c *fiber.Ctx) error {
	postUpdate := new(entities.Post)
	id := c.Query("id")
	
	if err := c.BodyParser(postUpdate); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if err := p.postUsecase.PostUsecaseUpdate(id, postUpdate); err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (p *postHandler) DeletePost(c *fiber.Ctx) error {
	id := c.Query("id")

	if err := p.postUsecase.PostUsecaseDelete(id); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusOK)
}