package controller

import (
	"alta/project/database"
	"alta/project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func CreateShoppingCartController(c echo.Context) error {
	// checking user id
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id user",
		})
	}
	// checking product id
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id product",
		})
	}
	// select product by id_product
	get_product, err := database.GetOneProduct(product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data product by id",
		})
	}
	//Binding JSON
	var shopping_carts model.Shopping_cart
	c.Bind(&shopping_carts)
	// insert product into cart
	new_shopping_cart, err := database.InsertProductIntoCart(user_id, product_id, get_product, shopping_carts.Qty)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":        "success create new shopping cart",
		"shopping cart": new_shopping_cart,
	})
}

func DeleteShoppingCartController(c echo.Context) error {
	// checking user_id
	user_id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid user id",
		})
	}
	// checking product_id
	product_id, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid product id",
		})
	}
	// canceling order by user id and product id
	delete_shopping_cart, err := database.DeleteShoppingCart(user_id, product_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "can not fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete shopiing cart",
		"users":  delete_shopping_cart,
	})
}
