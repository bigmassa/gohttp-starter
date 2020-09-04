package models

import (
	"gorm.io/gorm"
)

var Db *gorm.DB

func AutoMigrate() {
	Db.AutoMigrate(
		&User{},
	)
}

func SetDatabase(db *gorm.DB) {
	Db = db
}
