package models

type User struct {
	ID       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"varchar(300)" json:"username"`
	Password string `gorm:"varchar(300)" json:"password"`
	Role     string `gorm:"varchar(25)" json:"role"`
	Token    string `gorm:"varchar(300)" json:"token"`
}
