package controller

import (
	"alta/project/config"
	"alta/project/database"
	"alta/project/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
)

func testGetProductController(t *testing.T, controller echo.HandlerFunc) {
	// coba request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	controller(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var product_test []model.Product_test
	if err := json.Unmarshal(body, &product_test); err != nil {
		t.Error(err)
	}
	if len(product_test) != 1 {
		t.Errorf("expected one book, got: %#v", product_test)
		return
	}
	if product_test[0].Name != "Asus" {
		t.Errorf("expected Asus, got: %#v", product_test[0].Name)
	}
}

func TestGetProductController(t *testing.T) {
	// bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Product_test{})
	db.Delete(&model.Product_test{}, "1=1")
	m := model.NewGormProductModel(db)
	// bikin controller
	controller := GetManyProductsControllerTest(m)
	if err != nil {
		t.Error(err)
	}
	// insert data baru
	m.Insert(model.Product_test{
		Id_tes: 1,
		Name:   "Asus",
	})
	// test controller
	testGetProductController(t, controller)
	db.Delete(&model.Product_test{}, "1=1")
}
