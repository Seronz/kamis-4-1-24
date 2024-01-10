package order

import (
	"time"

	products "github.com/seronz/api/src/models/Products"
	"github.com/shopspring/decimal"
)

type Items struct {
	ID              string `json:"id" gorm:"size:36;uniqueIndex;not null;primary_key"`
	Order           Order
	OrderID         string `json:"order_id" gorm:"size:36;index"`
	Products        products.Products
	ProductsID      string          `json:"products_id" gorm:"size:36;index"`
	Qty             int             `json:"quantity"`
	BasePrice       decimal.Decimal `json:"base_price" gorm:"type:decimal(16,2)"`
	BaseTotal       decimal.Decimal `json:"base_total" gorm:"type:decimal(16,2)"`
	TaxAmount       decimal.Decimal `json:"tax_amount" gorm:"type:decimal(16,2)"`
	TaxPercent      decimal.Decimal `json:"tax_percent" gorm:"type:decimal(10,2)"`
	DiscountAmount  decimal.Decimal `json:"discount_amount" gorm:"type:decimal(16,2)"`
	DiscountPercent decimal.Decimal `json:"discount_percent" gorm:"type:decimal(10,2)"`
	SubTotal        decimal.Decimal `json:"sub_total" gorm:"type:decimal(16,2)"`
	Sku             string          `json:"sku" gorm:"size:36;index"`
	Name            string          `json:"name" gorm:"size:255"`
	Weight          decimal.Decimal `json:"weight" gorm:"type:decimal(10,2)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
