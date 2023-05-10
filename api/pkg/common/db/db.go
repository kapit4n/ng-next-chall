package db

import (
	"log"

	"github.com/kapit4n/ng-next-chall/api/pkg/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// have a flag to just migrate when it is required
	db.AutoMigrate(&models.Subject{})
	db.AutoMigrate(&models.Category{})

	return db
}
