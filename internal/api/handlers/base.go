package handlers

import (
	"github.com/bencoderus/auth-service/pkg/utils/http"
	"github.com/gofiber/fiber/v2"
)

func Home(context *fiber.Ctx) error {
	response := http.ApiResponse{StatusCode: 200, Message: "Auth service 1.0 is live"}

	return response.Send(context)
}
