package model

import "gorm.io/gorm"

type Shopping_cart struct {
	gorm.Model
	User_id    int    `json:"userid" form:"userid"`
	Product_id int    `json:"productid" form:"productid"`
	Name       string `json:"name" form:"name"`
	Category   string `json:"category" form:"category"`
	Type       string `json:"type" form:"type"`
	Price      int    `json:"price" form:"price"`
	Qty        int    `json:"qty" form:"qty"`
}
