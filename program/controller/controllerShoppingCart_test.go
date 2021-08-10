package controller

import (
	"alta/project/config"
	"alta/project/database"
	"alta/project/model"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func testPostShoppingCart(t *testing.T, controller echo.HandlerFunc) {
	// coba request
	reqBody, err := json.Marshal(map[string]int{
		"qty": 2,
	})
	if err != nil {
		t.Error(err)
		return
	}
	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	c.SetPath("/carts/:product_id")
	c.SetParamNames("product_id")
	c.SetParamValues("1")
	controller(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var cart model.Shopping_cart_test
	if err := json.Unmarshal(body, &cart); err != nil {
		t.Error(err)
	}
	if cart.Qty != 2 {
		t.Errorf("expected 2, got: %#v", cart.Qty)
	}
}

func TestPostShoppingCart(t *testing.T) {
	// bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Shopping_cart{})
	db.AutoMigrate(&model.Product_test{})
	db.Delete(&model.Shopping_cart_test{}, "1=1")
	db.Delete(&model.Product_test{}, "1=1")
	m := model.NewGormCartModel(db)
	n := model.NewGormProductModel(db)
	// bikin controller
	cartController := CreateShoppingCartControllerTest(m, n)
	if err != nil {
		t.Error(err)
	}
	n.Insert(model.Product_test{
		Id_tes: 1,
		Name:   "Asus",
	})
	// test controller
	testPostShoppingCart(t, cartController)
	db.Delete(&model.Shopping_cart_test{}, "1=1")
	db.Delete(&model.Product_test{}, "1=1")
}
