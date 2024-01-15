package users

import (
	"fmt"
	"time"

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

func CreateAddress(db *gorm.DB, address Address, user_token string) (Address, error) {
	fmt.Println(user_token)
	claims, err := jwt.JWTGetClaims(user_token)
	if err != nil {
		return Address{}, err
	}

	mycalims := claims.(struct {
		Sub        string
		Id         string
		Email      string
		Firstname  string
		Lastname   string
		Rememberme bool
		Userrole   string
	})

	fmt.Println("ini emailmu ", mycalims.Id)

	return Address{}, nil
}
