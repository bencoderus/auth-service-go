package main

import (
	"fmt"

	"github.com/bencoderus/auth-service/internal/api"
	"github.com/bencoderus/auth-service/internal/database"
	"github.com/bencoderus/auth-service/internal/database/models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
	}

	api.HandleRouting(app)

	db := database.ConnectToDB()
	db.AutoMigrate(&models.User{})

	app.Listen(":7700")
}
