package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/AlehBelski/go-card-api/controller/mocks"

	"github.com/stretchr/testify/assert"

	"github.com/AlehBelski/go-card-api/model"
)

func TestCartController_HandleCreate(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.NewCart(1, []model.CartItem{})

	expectedBodeResponse := `{"id":1,"items":[]}`

	serviceMock := new(mocks.Service)
	serviceMock.On("Create").Return(expectedCart, nil)

	c := NewCartController(serviceMock)

	req, err := http.NewRequest(http.MethodPost, "/carts", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	err = c.HandleCreate(rr, req)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(http.StatusOK, rr.Code, "handler returned wrong status code")

	actualBody := strings.TrimSuffix(rr.Body.String(), "\n")

	assertion.EqualValues(expectedBodeResponse, actualBody)
}

func TestCartController_HandleRead(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.NewCart(123, []model.CartItem{})

	expectedBodeResponse := `{"id":123,"items":[]}`

	serviceMock := new(mocks.Service)
	serviceMock.On("Read", 123).Return(expectedCart, nil)

	c := NewCartController(serviceMock)

	req, err := http.NewRequest(http.MethodGet, "/carts/123", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	err = c.HandleRead(rr, req)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(http.StatusOK, rr.Code, "handler returned wrong status code")

	actualBody := strings.TrimSuffix(rr.Body.String(), "\n")

	assertion.EqualValues(expectedBodeResponse, actualBody)
}

func TestCartController_HandleUpdate(t *testing.T) {
	assertion := assert.New(t)

	itemToUpdate := model.CartItem{}

	itemToUpdate.SetProduct("Shoes")
	itemToUpdate.SetQuantity(10)

	expectedCartItem := model.NewCartItem(1, 123, "Shoes", 10)

	expectedBodeResponse := `{"id":1,"cart_id":123,"product":"Shoes","quantity":10}`

	serviceMock := new(mocks.Service)
	serviceMock.On("Update", 123, itemToUpdate).Return(expectedCartItem, nil)

	c := NewCartController(serviceMock)

	req, err := http.NewRequest(http.MethodPost, "/carts/123/items", strings.NewReader(`{"Product":"Shoes","Quantity":10}`))

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	err = c.HandleUpdate(rr, req)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(http.StatusOK, rr.Code, "handler returned wrong status code")

	actualBody := strings.TrimSuffix(rr.Body.String(), "\n")

	assertion.EqualValues(expectedBodeResponse, actualBody)
}

func TestCartController_HandleRemove(t *testing.T) {
	assertion := assert.New(t)

	serviceMock := new(mocks.Service)
	serviceMock.On("DeleteItem", 123, 1).Return(nil)

	c := NewCartController(serviceMock)

	req, err := http.NewRequest(http.MethodDelete, "/carts/123/items/1", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	err = c.HandleRemove(rr, req)

	if err != nil {
		t.Fatal(err)
	}

	assertion.EqualValues(http.StatusOK, rr.Code, "handler returned wrong status code")
}
