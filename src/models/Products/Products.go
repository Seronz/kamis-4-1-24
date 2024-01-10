package products

import (
	"time"

	users "github.com/seronz/api/src/models/Users"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Products struct {
	ID               string `json:"id" gorm:"size:36;uniqueIndex;not null;primary_key"`
	ParentID         string `json:"parent_id" gorm:"size:36;index"`
	User             users.User
	UserID           string          `json:"user_id" gorm:"size:36;index"`
	ProductImage     []Image         `gorm:"foreignKey:product_id"`
	Categories       []Categories    `gorm:"many2many;product_categories"`
	Sku              string          `json:"sku" gorm:"size:100;index"`
	Name             string          `json:"name" gorm:"size:255"`
	Slug             string          `json:"slug" gorm:"size:255"`
	Price            decimal.Decimal `json:"price" gorm:"type:decimal(16,2)"`
	Stock            int             `json:"stock"`
	Weight           decimal.Decimal `json:"weight" gorm:"decimal(10,2)"`
	ShortDescription string          `json:"short_description" gorm:"size:255"`
	Description      string          `json:"description" gorm:"type:text"`
	Status           int             `json:"status" gorm:"default:0"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
