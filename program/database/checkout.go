package database

import (
	"alta/project/config"
	"alta/project/model"
)

//GET total price and Qty
func CheckoutRecord(id int) (model.Checkout, error) {
	var shopping_carts []model.Shopping_cart
	var checkout model.Checkout
	var user model.User
	if err := config.DB.Find(&user, "id = ?", id).Error; err != nil {
		return checkout, err
	}
	if err := config.DB.Find(&shopping_carts).Error; err != nil {
		return checkout, err
	}
	total_qty := 0
	total_price := 0
	for _, v := range shopping_carts {
		total_qty += v.Qty
		total_price = total_price + (v.Qty * v.Price)
	}
	checkout = model.Checkout{
		Total_qty:   total_qty,
		Total_price: total_price,
		User:        user.Name,
		User_id:     id,
	}
	if err := config.DB.Save(&checkout).Error; err != nil {
		return checkout, err
	}
	return checkout, nil
}
