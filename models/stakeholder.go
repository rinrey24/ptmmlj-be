package models

type Stakeholder struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	Company   string `json:"company"`
	Since     string `json:"since"`
	Is_active string `json:"is_active"`
}
