package http

import (
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	StatusCode int
	Message    string
	Data       any
	Error      any
}

type RestResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func (response ApiResponse) Send(c *fiber.Ctx) error {
	return c.Status(response.StatusCode).JSON(response.Json())
}

func (response ApiResponse) Json() RestResponse {
	var body RestResponse

	body.Status = response.StatusCode >= 200 && response.StatusCode <= 206
	body.Message = response.Message

	if response.Data != nil {
		body.Data = response.Data
	}

	if response.Error != nil {
		body.Error = response.Error
	}

	return body
}
