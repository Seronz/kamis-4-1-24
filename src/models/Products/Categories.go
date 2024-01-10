package products

import "time"

type Categories struct {
	ID        string `json:"id" gorm:"size:36;primary_key;not null;uniqueIndex"`
	ParentId  string `json:"parent_id" gorm:"size:36"`
	Section   Section
	SectionID string     `json:"section_id" gorm:"size:36;index"`
	Products  []Products `gorm:"many2many;product_categories"`
	Name      string     `json:"name" gorm:"size:100"`
	Slug      string     `json:"slug" gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
