package order

import (
	"time"

	users "github.com/seronz/api/src/models/Users"
)

type Customer struct {
	ID         string `json:"id" gorm:"size:36;uniqueIndex;not null;primary_key"`
	User       users.User
	UserID     string `json:"user_id" gorm:"size:36;index"`
	Order      Order
	OrderId    string `json:"order_id" gorm:"size:36;index"`
	FirstName  string `json:"first_name" gorm:"size:100;not_null"`
	LastName   string `json:"last_name" gorm:"size:100;not_null"`
	CityID     string `json:"city_id" gorm:"size:100"`
	ProvinceID string `json:"province_id" gorm:"size:100"`
	Address1   string `json:"address_1" gorm:"size:100"`
	Address2   string `json:"address_2" gorm:"size:100"`
	Phone      string `json:"phone" gorm:"size:50"`
	Email      string `json:"email" gorm:"size:100"`
	PostCode   string `json:"size:100" gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
