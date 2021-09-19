
package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/DaffaBarin/simple-books-mvc/models"
)

var DB *gorm.DB

func InitDB() {
	connectionString := "root:barin12345@tcp(localhost:3306)/Books?parseTime=True"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	MigrateDB()
	if err != nil {
		panic(err)
	}
}

func MigrateDB() {
	DB.AutoMigrate(&models.Book{})
}