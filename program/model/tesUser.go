package model

import "gorm.io/gorm"

type User_test struct {
	gorm.Model
	Id_tes   int
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Address  string `json:"address" form:"address"`
	Token    string `json :"token" form:"token"`
}
