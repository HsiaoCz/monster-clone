package db

import "gorm.io/gorm"

var instanceDB *gorm.DB

func GetDB() *gorm.DB {
	return instanceDB
}
