package models

import "time"

type History struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Date      string    `json:"date"`
	Event     string    `json:"event"`
	Is_active string    `json:"is_active"`
	Order     int       `json:"order"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy int       `json:"createdBy"`
	UpdatedBy int       `json:"updatedBy"`
}
