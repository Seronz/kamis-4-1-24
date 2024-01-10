package order

import (
	"time"

	users "github.com/seronz/api/src/models/Users"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	ID                  string `json:"id" gorm:"size:36;uniqueIndex;not null;primary_key"`
	User                users.User
	UserID              string `json:"user_id" gorm:"size:36;index"`
	Items               []Items
	Customer            *Customer
	Code                string `json:"code" gorm:"size:50;index"`
	Status              int
	OrderDate           time.Time
	PaymentDue          time.Time
	PaymentStatus       string          `json:"payment_status" gorm:"size:50;index"`
	PaymentToken        string          `json:"payment_token" gorm:"size:100;index"`
	BaseTotalPrice      decimal.Decimal `json:"base_total_price" gorm:"type:decimal(16,2)"`
	TaxAmount           decimal.Decimal `json:"tax_amount" gorm:"type:decimal(16,2)"`
	TaxPercent          decimal.Decimal `json:"percent" gorm:"type:decimal(10,2)"`
	DiscountAmount      decimal.Decimal `json:"discount_amount" gorm:"type:decimal(16,2)"`
	DiscountPercent     decimal.Decimal `json:"discount_percent" gorm:"type:decimal(10,2)"`
	ShippingCost        decimal.Decimal `json:"shipping_cost" gorm:"type:decimal(16,2)"`
	GrandTotal          decimal.Decimal `json:"grand_total" gorm:"type:decimal(16,2)"`
	Note                string          `json:"note" gorm:"type:text"`
	ShippingCourir      string          `json:"shipping_courir" gorm:"size:100"`
	ShippingServiceName string          `json:"shipping_service_name" gorm:"size:100"`
	ApproveBy           string          `json:"approve_by" gorm:"size:36"`
	ApprovedAt          time.Time
	CancledBy           string `json:"canceled_by" gorm:"size:36"`
	CancledAt           time.Time
	CancleationNote     string `json:"cancle_note" gorm:"size:255"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}
