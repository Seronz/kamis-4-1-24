package products

import "time"

type Image struct {
	ID         string `json:"id" gorm:"size:36;primary_key;not null;uniqueIndex"`
	Product    Products
	ProductID  string `json:"product_id" gorm:"size:36;index"`
	Path       string `json:"path" gorm:"type:text"`
	ExtraLarge string `json:"extra_large" gorm:"type:text;unique"`
	Large      string `json:"large" gorm:"type:text;unique"`
	Medium     string `json:"medium" gorm:"type:text;unique"`
	Small      string `json:"small" gorm:"type:text;unique"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
