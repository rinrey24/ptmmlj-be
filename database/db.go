package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/ptmmlj?parseTime=true"))
	if err != nil {
		fmt.Println("Gagal koneksi database")
	}

	// db.AutoMigrate(&models.User{})
	// db.AutoMigrate(&models.Article{})
	// db.AutoMigrate(&models.History{})
	// db.AutoMigrate(&models.Profile{})
	// db.AutoMigrate(&models.Stakeholder{})

	DB = db
}
