package model

import "gorm.io/gorm"

type Product_test struct {
	gorm.Model
	Id_tes int
	Name   string `json:"name" form:"name"`
}
type ProductModel interface {
	Get() []Product_test
	Insert(Product_test) error
}

type GormProductModel struct {
	db *gorm.DB
}

func (m *GormProductModel) Get() []Product_test {
	var carts []Product_test
	m.db.Find(&carts)
	return carts
}

func (m *GormProductModel) Insert(cart Product_test) error {
	tx := m.db.Save(&cart)
	return tx.Error
}

func NewGormProductModel(db *gorm.DB) *GormProductModel {
	return &GormProductModel{db: db}
}
