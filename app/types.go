package app

import (
	"time"
)

type DBConfig struct {
	DBDriver string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
}

type Cart struct {
	KodeProduk string    `json:"kodeProduk" gorm:"not null;size:191;primaryKey;autoIncrement:false;column:kodeProduk" form:"requestor" binding:"required"`
	NamaProduk string    `json:"namaProduk"  gorm:"not null;column:namaProduk"`
	Kuantitas  int       `json:"kuantitas" gorm:"not null;column:kuantitas"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type AddProductRequest struct {
	KodeProduk string `json:"kodeProduk" form:"kodeProduk" binding:"required"`
	NamaProduk string `json:"namaProduk" form:"namaProduk" binding:"required"`
	Kuantitas  int    `json:"kuantitas" form:"kuantitas" binding:"required"`
}

type DeleteProductRequest struct {
	KodeProduk string `json:"kodeProduk" form:"kodeProduk" binding:"required"`
}

type ShowCartRequest struct {
	NamaProduk string `json:"namaProduk" form:"namaProduk"`
	Kuantitas  string `json:"kuantitas" form:"kuantitas"`
}
