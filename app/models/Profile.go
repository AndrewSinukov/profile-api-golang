package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Profile struct {
	gorm.Model
	Codes int `gorm:"column:codes"`
	Name string `gorm:"column:name" json:"name"`
	Author string   `gorm:"column:author" json:"author"`
	Category string `gorm:"column:category" json:"category"`
}

func (b *Profile) TableName() string {
  return "Profiles"
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
	if err := db.Where("codes = ?", ids).First(&b).Error; err != nil {
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