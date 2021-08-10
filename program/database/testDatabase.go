package database

import (
	"alta/project/config"
	"alta/project/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDB(connectionString string) (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return DB, err
	}
	return DB, err
}

func InsertProductIntoCartTes(id_product int, qty int, cartModel model.CartModel) (model.Shopping_cart_test, error) {
	// insert and select data product to cart
	shopping_cart_test := model.Shopping_cart_test{
		Product_id: id_product,
		Name:       "Asus",
		Qty:        qty,
	}
	cartModel.Insert(shopping_cart_test)
	return shopping_cart_test, nil
}

func GetProductTest() (interface{}, error) {
	var products []model.Product_test
	if err := config.DB.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}
