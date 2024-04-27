package models

import (
	"errors"
	"time"

	"github.com/bencoderus/auth-service/internal/database"
	"github.com/bencoderus/auth-service/internal/types"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id,omitempty" gorm:"primaryKey"`
	Name      string         `json:"name,omitempty"`
	Email     string         `json:"email,omitempty" gorm:"uniqueIndex"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deletedAt,omitempty"`
}

func GetUserByEmail(email string) User {
	db := database.GetDBConnection()
	user := &User{}

	db.First(user, "email = ?", email)

	return *user
}

func GetUserById(id float64) User {
	db := database.GetDBConnection()
	user := &User{}

	db.First(user, "id = ?", id)

	return *user
}

func CreateUser(payload types.RegisterPayload) (User, error) {
	db := database.GetDBConnection()
	user := &User{Name: payload.Name, Email: payload.Email, Password: payload.Password}
	db.Create(user)

	if user.ID == 0 {
		return *user, errors.New("unable to create user")
	}

	return *user, nil
}
