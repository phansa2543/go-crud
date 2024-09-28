package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type PostHandler interface {
	CreatePost(c *fiber.Ctx) error
	ReadById(c *fiber.Ctx) error
	ReadAllTitle(c *fiber.Ctx) error
	UpdatePost(c *fiber.Ctx) error
	DeletePost(c *fiber.Ctx) error
}