package routes

import (
	"alta/project/constant"
	"alta/project/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	//Basic Routes for Skeleton **Dont Touch**
	e.GET("/users", controller.GetManyUsersController)
	e.GET("/products", controller.GetManyProductsController)
	e.POST("/products", controller.NewItemController)
	e.DELETE("/users/:id", controller.DeleteUserController)

	//Work your code here

	//GET show all product by category and type (Mba patmiza)
	e.GET("/products/:category", controller.GetProductCategoryController)
	e.GET("/products/:category/:type", controller.GetProductTypeController)

	//POST Add product to shopping cart (Mas Doni)
	e.POST("carts/:user_id/:product_id", controller.CreateShoppingCartController)
	//DELETE product to shopping cart (Mas Doni)
	e.DELETE("carts/:user_id/:product_id", controller.DeleteShoppingCartController)

	//Fawwaz Workspace
	//GET product form shopping cart
	e.GET("/cart", controller.ProductInCart)

	//Register User
	e.POST("/users", controller.UserRegister)

	//Login
	e.POST("/login", controller.LoginUser)

	//Checkout
	eJwt := e.Group("/jwt")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	eJwt.POST("/checkout/:id", controller.Checkout)

	//Transaction
	eJwt.POST("/transaction/:id", controller.TransactionController)

}
