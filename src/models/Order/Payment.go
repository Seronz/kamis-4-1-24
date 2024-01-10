package order

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Payment struct {
	ID          string `json:"id" gorm:"size:36;not null;uniqueIndex;primary_key"`
	Order       Order
	OrderID     string          `json:"order_id" gorm:"size:36;index"`
	Number      string          `json:"number" gorm:"size:100;index"`
	Amount      decimal.Decimal `json:"amount" gorm:"type:decimal(16,2)"`
	Method      string          `json:"method" gorm:"size:100"`
	Status      string          `json:"status" gorm:"size:100"`
	Token       string          `json:"token" gorm:"size:100;index"`
	Payload     string          `json:"payload" gorm:"type:text"`
	PaymentType string          `json:"payment_type" gorm:"size:100"`
	VaNumber    string          `json:"va_number" gorm:"size:100"`
	BillCode    string          `json:"bill_code" gorm:"size:100"`
	BillKey     string          `json:"bill_key" gorm:"size:100"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
