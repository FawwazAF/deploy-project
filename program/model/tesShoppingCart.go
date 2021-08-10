package model

import "gorm.io/gorm"

type Shopping_cart_test struct {
	gorm.Model
	Product_id int    `json:"productid" form:"productid"`
	Name       string `json:"name" form:"name"`
	Qty        int    `json:"qty" form:"qty"`
}

type CartModel interface {
	Get() []Shopping_cart_test
	Insert(Shopping_cart_test) error
}

type GormCartModel struct {
	db *gorm.DB
}

func (m *GormCartModel) Get() []Shopping_cart_test {
	var carts []Shopping_cart_test
	m.db.Find(&carts)
	return carts
}

func (m *GormCartModel) Insert(cart Shopping_cart_test) error {
	tx := m.db.Save(&cart)
	return tx.Error
}

func NewGormCartModel(db *gorm.DB) *GormCartModel {
	return &GormCartModel{db: db}
}
