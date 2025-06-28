package api

import (
	"task-workmate/internal/service"
	
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Routers struct {
	Service service.Service
}

func NewRouters(r *Routers) *fiber.App {
	app := fiber.New()

	// Настройка CORS
	app.Use(cors.New(cors.Config{
		AllowMethods:  "GET, POST, DELETE",
		AllowHeaders:  "Accept, Authorization, Content-Type, X-CSRF-Token, X-REQUEST-ID",
		ExposeHeaders: "Link",
		MaxAge:        300,
	}))

	// Группа маршрутов
	apiGroup := app.Group("/tasks")
	apiGroup.Post("/", r.Service.CreateTask)
	apiGroup.Delete("/:id", r.Service.DeleteTask)
	apiGroup.Get("/:id", r.Service.GetTask)

	return app
}
