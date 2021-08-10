package database

import (
	"alta/project/config"
	"alta/project/model"
)

//GET total price and Qty
func Transaction(checkout model.Checkout) (interface{}, error) {
	transaction := model.Transaction{
		Status:            "success",
		User:              checkout.User,
		Total_transaction: checkout.Total_price,
	}
	if err := config.DB.Save(&transaction).Error; err != nil {
		return nil, err
	}
	return transaction, nil
}

func GetCheckout(user_id int) (model.Checkout, error) {
	var checkout model.Checkout
	if err := config.DB.Find(&checkout, "user_id=?", user_id).Error; err != nil {
		return checkout, err
	}
	return checkout, nil
}

func CleanAfter(id int) error {
	var shopping_carts []model.Shopping_cart
	// if err := config.DB.Find(&shopping_carts, "user_id=?", id).Error; err != nil {
	// 	return err
	// }
	if err := config.DB.Delete(&shopping_carts, "user_id=?", id).Error; err != nil {
		return err
	}
	var checkout model.Checkout
	if err := config.DB.Delete(&checkout, "user_id=?", id).Error; err != nil {
		return err
	}
	return nil
}
