package models

import "time"

type Profile struct {
	ID           int64     `gorm:"primaryKey" json:"id"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Slogan       string    `json:"slogan"`
	Description  string    `json:"description"`
	Telephone    string    `json:"telephone"`
	Mobile_phone string    `json:"mobile_phone"`
	Email        string    `json:"email"`
	UpdatedAt    time.Time `json:"updatedAt"`
	UpdatedBy    int       `json:"updatedBy"`
}
