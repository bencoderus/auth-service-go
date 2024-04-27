package handlers

import (
	"github.com/bencoderus/auth-service/internal/services"
	"github.com/bencoderus/auth-service/pkg/utils/http"
	"github.com/gofiber/fiber/v2"
)

func GetUserProfile(context *fiber.Ctx) error {
	value := context.Locals("userId")
	serviceUnavailable := http.ApiResponse{StatusCode: 503, Message: "Unable to fetch user profile"}

	if value == nil {
		return serviceUnavailable.Send(context)
	}

	userId := value.(float64)
	user, err := services.GetProfile(userId)

	if err != nil {
		return serviceUnavailable.Send(context)
	}

	response := http.ApiResponse{StatusCode: 200, Message: "User profile successful", Data: user}

	return response.Send(context)
}
