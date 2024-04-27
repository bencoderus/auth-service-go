package services

import (
	"errors"

	"github.com/bencoderus/auth-service/internal/database/models"
	"github.com/bencoderus/auth-service/internal/types"
)

type AuthResponse struct {
	User  models.User `json:"user"`
	Token JwtToken    `json:"token"`
}

func GetProfile(userId float64) (models.User, error) {
	user := models.GetUserById(userId)

	if user.ID == 0 {
		return user, errors.New("profile not found")
	}

	return user, nil
}

func Login(payload types.LoginPayload) (AuthResponse, error) {
	user := models.GetUserByEmail(payload.Email)

	if user.ID == 0 {
		return AuthResponse{}, errors.New("email is invalid")
	}

	valid := VerifyPasswordHash(user.Password, payload.Password)

	if !valid {
		return AuthResponse{}, errors.New("password is invalid")
	}

	token, err := SignToken(user.ID)

	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		User:  user,
		Token: token,
	}, err
}

func CreateUser(payload types.RegisterPayload) (AuthResponse, error) {
	found := models.GetUserByEmail(payload.Email)

	if found.ID != 0 {
		return AuthResponse{}, errors.New("user already exists")
	}

	passwordHashByte, err := HashPassword(payload.Password)

	if err != nil {
		return AuthResponse{}, err
	}

	payload.Password = string(passwordHashByte)
	user, err := models.CreateUser(payload)

	if err != nil {
		return AuthResponse{}, err
	}

	token, err := SignToken(user.ID)

	if err != nil {
		return AuthResponse{}, err
	}

	return AuthResponse{
		User:  user,
		Token: token,
	}, err
}
