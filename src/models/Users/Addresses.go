package users

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	jwt "github.com/seronz/api/src/utils/JWT"
	"gorm.io/gorm"
)

type Address struct {
	ID         string `gorm:"size:36;not null;primary_key;uniqueIndex" json:"id"`
	User       User
	UserID     string `gorm:"size:36;index"`
	Name       string `gorm:"size:255;not null" json:"name"`
	IsPrimary  bool
	CityID     string `gorm:"size:100"`
	ProvinceID string `gorm:"size:100"`
	Address1   string `gorm:"size:100"`
	Address2   string `gorm:"size:100"`
	Phone      string `gorm:"size:15"`
	Email      string `gorm:"size:100"`
	PostCode   string `gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (l *Address) CreateAddress(db *gorm.DB) error {
	result := db.Create(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (l *Address) GetAddress(db *gorm.DB) error {
	result := db.Select("*").Where("id = ?", l.ID).First(&l)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateAddress(db *gorm.DB, user_token string) (Address, error) {
	claims, err := jwt.JWTGetClaims(user_token)
	if err != nil {
		return Address{}, err
	}

	var a Address

	a.ID = uuid.NewString()
	a.UserID = claims.ID
	a.Name = claims.Firstname + claims.Lastname
	fmt.Printf("ini id mu : %s \n", claims.ID)
	// err = a.CreateAddress(db)
	// if err != nil {
	// 	return Address{}, err
	// }

	return Address{}, nil
}
