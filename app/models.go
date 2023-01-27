package app

import (
	"sort"
	"strings"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartModel interface {
	Init()
	AddProduct(req AddProductRequest) (res Cart, err error)
	DeleteProduct(req DeleteProductRequest) (res bool, err error)
	ShowCart(req ShowCartRequest) (res []Cart, err error)
}

type cartModel struct {
	db *gorm.DB
}

func NewCartModels() CartModel {
	return &cartModel{}
}

func (m *cartModel) Init() {
	db := NewSQL().SetupDatabaseConnection()
	m.db = db //calling database connection
}

func (m *cartModel) AddProduct(req AddProductRequest) (res Cart, err error) {
	cart := Cart{
		KodeProduk: req.KodeProduk,
		NamaProduk: req.NamaProduk,
		Kuantitas:  req.Kuantitas,
	}
	tx := m.db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "kodeProduk"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"kuantitas": gorm.Expr("kuantitas + ?", req.Kuantitas), "namaProduk": req.NamaProduk}),
	}).Create(&cart)

	if tx.Error != nil {
		return Cart{}, tx.Error
	}
	if tx.RowsAffected == 0 { //check if request beetwen user still pending
		return cart, gorm.ErrRecordNotFound
	}
	tx.Find(&cart)

	return cart, err
}

func (m *cartModel) DeleteProduct(req DeleteProductRequest) (res bool, err error) {

	cart := Cart{}
	tx1 := m.db.Debug().Where("kodeProduk", req.KodeProduk).Delete(&cart)
	if tx1.Error != nil { //check if request beetwen user still pending
		return false, tx1.Error
	}
	if tx1.RowsAffected == 0 {
		return false, gorm.ErrRecordNotFound
	}

	return true, err
}

func (m *cartModel) ShowCart(req ShowCartRequest) (res []Cart, err error) {
	tx1 := m.db.Debug().Where("namaProduk LIKE ?", "%"+req.NamaProduk+"%").Find(&res)
	if tx1.Error != nil {
		return res, tx1.Error
	}
	if tx1.RowsAffected == 0 {
		return res, gorm.ErrRecordNotFound

	}

	switch strings.ToLower(req.Kuantitas) {
	case "asc":
		sort.Slice(res, func(i, j int) bool {
			return res[i].Kuantitas < res[j].Kuantitas
		})
	case "desc":
		sort.Slice(res, func(i, j int) bool {
			return res[i].Kuantitas > res[j].Kuantitas
		})
	}

	return res, err
}
