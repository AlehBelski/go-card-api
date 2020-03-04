package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlehBelski/go-card-api/model"
	"github.com/stretchr/testify/mock"
)

type CartServiceMock struct {
	mock.Mock
}

func (srv *CartServiceMock) Create() (model.Cart, error) {
	args := srv.Called()

	return args.Get(0).(model.Cart), args.Error(1)
}

func (srv *CartServiceMock) Read(id int) (model.Cart, error) {
	args := srv.Called(id)

	return args.Get(0).(model.Cart), args.Error(1)
}

func (srv *CartServiceMock) Update(id int, item model.CartItem) (model.CartItem, error) {
	args := srv.Called(id, item)

	return args.Get(0).(model.CartItem), args.Error(1)
}

func (srv *CartServiceMock) DeleteItem(cartId, itemId int) error {
	args := srv.Called(cartId, itemId)

	return args.Error(0)
}

func TestCartController_HandleCreate(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.Cart{
		ID:    1,
		Items: []model.CartItem{},
	}

	expectedBodeResponse := `{"ID":1,"Items":[]}`

	serviceMock := new(CartServiceMock)
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

	expectedCart := model.Cart{
		ID:    123,
		Items: []model.CartItem{},
	}

	expectedBodeResponse := `{"ID":123,"Items":[]}`

	serviceMock := new(CartServiceMock)
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

	itemToUpdate := model.CartItem{
		Product:  "Shoes",
		Quantity: 10,
	}

	expectedCartItem := model.CartItem{
		ID:       1,
		CartID:   123,
		Product:  "Shoes",
		Quantity: 10,
	}

	expectedBodeResponse := `{"ID":1,"CartID":123,"Product":"Shoes","Quantity":10}`

	serviceMock := new(CartServiceMock)
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

	serviceMock := new(CartServiceMock)
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
