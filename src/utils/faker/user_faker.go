package faker

import (
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	users "github.com/seronz/api/src/models/Users"
	"gorm.io/gorm"
)

func UserFaker(db *gorm.DB) *users.User {
	return &users.User{
		Id:            uuid.New().String(),
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "password",
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeletedAt:     gorm.DeletedAt{},
	}
}
