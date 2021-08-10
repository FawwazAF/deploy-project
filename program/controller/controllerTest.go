package controller

import (
	"alta/project/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetManyProductsControllerTest(productModel model.ProductModel) echo.HandlerFunc {
	return func(c echo.Context) error {
		users := productModel.Get()
		return c.JSON(http.StatusOK, users)
	}
}

func CreateShoppingCartControllerTest(cartModel model.CartModel, productModel model.ProductModel) echo.HandlerFunc {
	return func(c echo.Context) error {

		product_id, err := strconv.Atoi(c.Param("product_id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, product_id)
		}
		product_get := productModel.Get()
		// Binding JSON
		var Shopping_cart_tes model.Shopping_cart_test
		c.Bind(&Shopping_cart_tes)
		shopping_cart_test := model.Shopping_cart_test{
			Product_id: product_id,
			Name:       product_get[0].Name,
			Qty:        Shopping_cart_tes.Qty,
		}
		cartModel.Insert(shopping_cart_test)
		return c.JSON(http.StatusOK, shopping_cart_test)
	}
}
