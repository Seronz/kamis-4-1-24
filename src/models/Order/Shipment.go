package order

import (
	"time"

	users "github.com/seronz/api/src/models/Users"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Shiment struct {
	ID          string `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        users.User
	UserID      string `json:"user_id" gorm:"size:36;index"`
	Order       Order
	OrderID     string `json:"order_id" gorm:"size:36;index"`
	TrackNumber string `json:"track_number" gorm:"size:255;index"`
	Status      string `json:"status" gorm:"size:36;index"`
	TotalQty    int
	TotalWeight decimal.Decimal `json:"total_weight" gorm:"type:decimal(10,2)"`
	FirstName   string          `json:"first_name" gorm:"size:100;not_null"`
	LastName    string          `json:"last_name" gorm:"size:100;not_null"`
	CityID      string          `json:"city_id" gorm:"size:100"`
	ProvinceID  string          `json:"province_id" gorm:"size:100"`
	Address1    string          `json:"address_1" gorm:"size:100"`
	Address2    string          `json:"address_2" gorm:"size:100"`
	Phone       string          `json:"phone" gorm:"size:50"`
	Email       string          `json:"email" gorm:"size:100"`
	PostCode    string          `json:"size:100" gorm:"size:100"`
	ShipedBy    string          `json:"shiped_by" gorm:"size:36"`
	ShipedAt    time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
