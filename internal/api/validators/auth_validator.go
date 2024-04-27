package validators

import (
	"fmt"
	"regexp"

	"github.com/bencoderus/auth-service/internal/types"
)

func emailIsValid(email string) bool {
	matched, _ := regexp.MatchString("^[^\\s@]+@[^\\s@]+\\.[^\\s@]+$", email)

	return matched
}

func ValidateRegisterPayload(payload types.RegisterPayload) error {
	if payload.Email == "" {
		return fmt.Errorf("email is required")
	}

	if !emailIsValid(payload.Email) {
		return fmt.Errorf("email is not valid")
	}

	if payload.Password == "" {
		return fmt.Errorf("password is required")
	}

	if payload.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(payload.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if len(payload.Name) < 3 {
		return fmt.Errorf("name must be at least 8 characters long")
	}

	return nil
}

func ValidateLoginPayload(payload types.LoginPayload) error {
	if payload.Email == "" {
		return fmt.Errorf("email is required")
	}

	if !emailIsValid(payload.Email) {
		return fmt.Errorf("email is not valid")
	}

	if payload.Password == "" {
		return fmt.Errorf("password is required")
	}

	if len(payload.Password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	return nil
}
