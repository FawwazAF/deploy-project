package model

import "gorm.io/gorm"

type Checkout struct {
	gorm.Model
	User        string
	User_id     int
	Total_qty   int
	Total_price int
}
