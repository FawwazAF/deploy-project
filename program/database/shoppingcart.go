package database

import (
	"alta/project/config"
	"alta/project/model"
)

// function to insert product into shopping cart
func InsertProductIntoCart(id_user, id_product int, product model.Product, qty int) (interface{}, error) {
	// insert and select data product to cart
	shoppingCart := model.Shopping_cart{
		User_id:    id_user,
		Product_id: id_product,
		Name:       product.Name,
		Category:   product.Category,
		Type:       product.Type,
		Price:      product.Price,
		Qty:        qty,
	}
	if err := config.DB.Save(&shoppingCart).Error; err != nil {
		return shoppingCart, err
	}
	return shoppingCart, nil
}

// function to cancel order or delete product from shopping cart
func DeleteShoppingCart(user_id, product_id int) (interface{}, error) {
	var cart model.Shopping_cart
	if err := config.DB.Where("user_id=? AND product_id=?", user_id, product_id).First(&cart).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

//GET Product in shopping_cart
func GetProductInCart() (interface{}, error) {
	var shopping_carts []model.Shopping_cart
	if err := config.DB.Find(&shopping_carts).Error; err != nil {
		return nil, err
	}
	return shopping_carts, nil
}

func GetOneProduct(product_id int) (model.Product, error) {
	var product model.Product
	if err := config.DB.Find(&product, "id=?", product_id).Error; err != nil {
		return product, err
	}
	return product, nil
}
