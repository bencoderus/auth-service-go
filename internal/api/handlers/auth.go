package handlers

import (
	"fmt"

	"github.com/bencoderus/auth-service/internal/api/validators"
	"github.com/bencoderus/auth-service/internal/services"
	"github.com/bencoderus/auth-service/internal/types"
	"github.com/bencoderus/auth-service/pkg/utils/http"
	"github.com/gofiber/fiber/v2"
)

func Login(context *fiber.Ctx) error {
	payload := http.ParseBody[types.LoginPayload](context.Body())

	error := validators.ValidateLoginPayload(payload)

	if error != nil {
		response := http.ApiResponse{StatusCode: 400, Message: "Validation error.", Error: error.Error()}

		return response.Send(context)
	}

	authResponse, err := services.Login(payload)

	if err != nil {
		fmt.Println(err)
		response := http.ApiResponse{StatusCode: 400, Message: "Invalid credentials."}

		return response.Send(context)
	}

	response := http.ApiResponse{StatusCode: 200, Message: "Login successful", Data: authResponse}

	return response.Send(context)
}

func Register(context *fiber.Ctx) error {
	payload := http.ParseBody[types.RegisterPayload](context.Body())

	err := validators.ValidateRegisterPayload(payload)

	if err != nil {
		response := http.ApiResponse{StatusCode: 400, Message: "Validation error.", Error: err.Error()}

		return response.Send(context)
	}

	authResponse, err := services.CreateUser(payload)

	if err != nil {
		response := http.ApiResponse{StatusCode: 400, Message: "Unable to create account.", Error: err.Error()}

		return response.Send(context)
	}

	response := http.ApiResponse{StatusCode: 201, Message: "User created successfully", Data: authResponse}

	return response.Send(context)
}
