package database

import (
	"github.com/lutfiharidha/haioo/app"
	"gorm.io/gorm"
)

func Migrator(db *gorm.DB) {
	db.AutoMigrate(&app.Cart{})
}
