package api

import (
	"github.com/bencoderus/auth-service/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func HandleRouting(app *fiber.App) {
	app.Get("/", handlers.Home)

	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)

	userRoute := app.Group("/user", AuthMiddleware)
	userRoute.Get("/profile", handlers.GetUserProfile)
}
