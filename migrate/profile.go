package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Profile struct {
	gorm.Model
	ID int `gorm:"column:id" json:"id"`
	FirstName string `gorm:"column:first_name" json:"first_name"`
	LastName string   `gorm:"column:last_name" json:"last_name"`
	PhoneNumber string `gorm:"column:phone_number" json:"phone_number"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Profile{})
	hasUser := db.HasTable(&Profile{})
	fmt.Println("Table profile is ", hasUser)
	if !hasUser {
		db.CreateTable(&Profile{})
	}

	return db
}
