package controller

import (
	"alta/project/database"
	_ "alta/project/middleware"
	"net/http"

	"github.com/labstack/echo"
)

//Work Your Code Here

func ProductInCart(c echo.Context) error {
	// user_id, err := strconv.Atoi(c.Param("userid"))
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "please login first",
	// 	})
	// }
	shopping_carts, err := database.GetProductInCart()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product in cart",
		"user":    shopping_carts,
	})
}

func GetProductCategoryController(c echo.Context) error {
	product_category := c.Param("category")
	GetProductCategory, err := database.GetProductCategory(product_category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product category",
		"product": GetProductCategory,
	})
}

func GetProductTypeController(c echo.Context) error {
	product_category := c.Param("category")
	product_type := c.Param("type")
	GetProductType, err := database.GetProductType(product_category, product_type)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product category",
		"product": GetProductType,
	})
}
