package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Status            string
	User              string
	Total_transaction int
	Paid              int `json:"paid" form:"paid"`
}
