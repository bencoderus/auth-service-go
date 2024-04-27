package api

import (
	"strings"

	"github.com/bencoderus/auth-service/internal/services"
	"github.com/bencoderus/auth-service/pkg/utils/http"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()

	authorizationHeader := headers["Authorization"]
	unauthenticatedResponse := http.ApiResponse{StatusCode: 401, Message: "Unauthenticated"}

	if len(authorizationHeader) == 0 {
		return unauthenticatedResponse.Send(c)
	}

	token := authorizationHeader[0]

	if token == "" {
		return unauthenticatedResponse.Send(c)
	}

	authToken := strings.Replace(token, "Bearer ", "", 1)

	jwtToken, err := services.ParseToken(authToken)

	if err != nil {
		return unauthenticatedResponse.Send(c)
	}

	tokenData := services.ExtractJwtTokenClaim(jwtToken)

	c.Locals("userId", tokenData["username"])

	return c.Next()
}
