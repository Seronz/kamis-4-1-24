package products

import "time"

type Section struct {
	ID         string `json:"id" gorm:"size:36;not null;primary_key;uniqueIndex"`
	Name       string `json:"name" gorm:"size:100"`
	Slug       string `json:"slug" gorm:"size:100"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Categories []Categories
}
