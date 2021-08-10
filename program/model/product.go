package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name"`
	Category    string `json:"category" form:"category"`
	Type        string `json:"type" form:"type"`
	Price       int    `json:"price" form:"price"`
	Description string `json:"description" form:"description"`
	Stock       int    `json:"stock" form:"stock"`
}
