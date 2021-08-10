package database

import (
	"alta/project/config"
	"alta/project/middleware"
	"alta/project/model"
	"errors"
)

func GetUsers() (interface{}, error) {
	var users []model.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteUser(id int) (interface{}, error) {
	var users []model.User
	if err := config.DB.Find(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	if err := config.DB.Delete(&users, "id=?", id).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CheckDuplicate(user model.User) error {
	var users model.User
	err := config.DB.Where("email = ? ", user.Email).First(&users).Error
	if err != nil {
		return nil
	}
	err = errors.New("email already exist")
	return err
}

func CreateUser(users model.User) (interface{}, error) {
	if err := config.DB.Save(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func LoginUsers(email, password string) (interface{}, error) {
	var user model.User
	var err error
	if err = config.DB.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}
	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, err
}
