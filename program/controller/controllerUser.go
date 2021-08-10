package controller

import (
	"alta/project/database"
	"alta/project/model"
	"net/http"

	"github.com/labstack/echo"
)

//Work Your Code Here

func UserRegister(c echo.Context) error {
	// binding data
	user := model.User{}
	c.Bind(&user)
	//check if email exist
	duplicate := database.CheckDuplicate(user)
	if duplicate != nil {
		return echo.NewHTTPError(http.StatusBadRequest, duplicate.Error())
	}
	users, err := database.CreateUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     users,
	})
}

func LoginUser(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	users, err := database.LoginUsers(user.Email, user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "status login",
		"users":  users,
	})
}
