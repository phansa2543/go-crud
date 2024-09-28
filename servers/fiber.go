package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/phansa2543/go-crud/config"
	"github.com/phansa2543/go-crud/databases"
	"github.com/phansa2543/go-crud/handlers"
	"github.com/phansa2543/go-crud/repositories"
	"github.com/phansa2543/go-crud/usecases"
)

type FiberServer struct {
	app *fiber.App
	db databases.Database
}

func NewFiberServer(db databases.Database) Server {
	app := fiber.New()
	return &FiberServer{
		app: app,
		db: db,
	}
}

func (f *FiberServer) Start() {
	conf := config.LoadConfig()

	f.app.Use(recover.New())
	f.app.Use(logger.New())

	f.app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).SendString("Hello")
	})

	v1 := f.app.Group("v1")

	f.postHandler(v1)

	f.app.Listen(fmt.Sprintf(":%s", conf.Server.Port))
}

func (f *FiberServer) postHandler(v1 fiber.Router) {
	postRepo := repositories.NewPostPostgresRepository(f.db)
	postUsecase := usecases.NewPostUsecaseImpl(postRepo)
	postHandler := handlers.NewPostHandler(postUsecase)

	postRoute := v1.Group("post")
	// Create Post
	postRoute.Post("", postHandler.CreatePost)
	// Get Post By Title
	postRoute.Get(":id", postHandler.ReadById)
	// Get All
	postRoute.Get("", postHandler.ReadAllTitle)
	// Update Post
	postRoute.Put("", postHandler.UpdatePost)
	// Delete Post
	postRoute.Delete("", postHandler.DeletePost)
}