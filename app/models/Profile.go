package models

import (
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

type Address struct {
	City  string `gorm:"column:city" json:"city"`
	State string `gorm:"column:state" json:"state"`
}

func (b *Profile) TableName() string {
  return "profiles"
}

func InsertProfile(db *gorm.DB, b *Profile) (err error) {
	if err = db.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProfile(db *gorm.DB, b *[]Profile) (err error) {
	if err = db.Order("id desc").Find(b).Error; err != nil {
		return err
	}
	return nil
}

func OneProfileGetting(db *gorm.DB, ids int, b *Profile) (err error) {
	if err := db.Where("id = ?", ids).First(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateProfile(db *gorm.DB, b *Profile) (err error) {
	if err = db.Save(b).Error; err != nil {
		return err
	}
	return nil
}

func DeletedProfile(db *gorm.DB, b *Profile) (err error) {
	if err = db.Delete(b).Error; err != nil {
		return err
	}
	return nil
}