package users

import (
	"time"
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
