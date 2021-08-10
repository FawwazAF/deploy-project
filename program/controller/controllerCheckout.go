package controller

import (
	"alta/project/database"
	"alta/project/middleware"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Checkout(c echo.Context) error {
	//Check Authentication
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	loggedInUserId := middleware.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access, you can only see your own")
	}

	// Checkout
	total, err := database.CheckoutRecord(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages":    "Checkout Success",
		"User":        total.User,
		"User ID":     total.User_id,
		"total_qty":   total.Total_qty,
		"total_price": total.Total_price,
	})
}

func GetCheckout(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "please login first",
		})
	}
	checkout, err := database.GetCheckout(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product in cart",
		"user":    checkout,
	})
}
