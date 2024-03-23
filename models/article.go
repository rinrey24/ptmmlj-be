package models

import "time"

type Article struct {
	ID          int64     `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"varchar(300)" json:"title"`
	Description string    `gorm:"varchar(max)" json:"description"`
	Slug        string    `gorm:"varchar(250)" json:"slug"`
	Is_active   string    `gorm:"integer" json:"is_active"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   time.Time `json:"createdBy"`
	UpdatedBy   time.Time `json:"updatedBy"`
}
