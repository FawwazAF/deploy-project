package controller

import (
	"alta/project/database"
	"alta/project/middleware"
	"alta/project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func TransactionController(c echo.Context) error {
	//Check Authentication
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	loggedInUserId := middleware.ExtractTokenUserId(c)
	if loggedInUserId != id {
		return echo.NewHTTPError(http.StatusUnauthorized, "unauthorized access, please login first")
	}

	//Fetch data from checkout
	checkout, err := database.GetCheckout(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	//Checking if payment valid
	var transaction model.Transaction
	c.Bind(&transaction)
	if checkout.Total_price == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "no fetched data")
	}
	if transaction.Paid < checkout.Total_price {
		return echo.NewHTTPError(http.StatusBadRequest, "transaction error,  not enough money")
	}

	// save history
	result, err := database.Transaction(checkout)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	//Clean Cart
	if err := database.CleanAfter(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"transaction": result,
	})
}
