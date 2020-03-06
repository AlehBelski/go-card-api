package service

import (
	"testing"

	"github.com/AlehBelski/go-card-api/model"
	"github.com/AlehBelski/go-card-api/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCartService_Create(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.Cart{
		ID:    123,
		Items: []model.CartItem{},
	}

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("Create").Return(expectedCart, nil)

	cartService := NewCartService(cartRepository)

	actualCart, err := cartService.Create()

	assertion.Nil(err)
	assertion.EqualValues(expectedCart, actualCart)

	cartRepository.AssertExpectations(t)
}

func TestCartService_Read(t *testing.T) {
	assertion := assert.New(t)

	ID := 123
	expectedCart := model.Cart{
		ID:    ID,
		Items: []model.CartItem{},
	}

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("IsCartExists", ID).Return(true, nil)
	cartRepository.On("Read", expectedCart.ID).Return(expectedCart, nil)

	cartService := NewCartService(cartRepository)

	actualCart, err := cartService.Read(expectedCart.ID)

	assertion.Nil(err)
	assertion.EqualValues(expectedCart, actualCart)

	cartRepository.AssertExpectations(t)
}

func TestCartService_Update(t *testing.T) {
	assertion := assert.New(t)

	ID := 123

	itemToUpdate := model.CartItem{
		Product:  "Shoes",
		Quantity: 100,
	}

	expectedCartItem := model.CartItem{
		CartID:   ID,
		Product:  itemToUpdate.Product,
		Quantity: itemToUpdate.Quantity,
	}

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("IsCartExists", ID).Return(true, nil)
	cartRepository.On("Update", ID, itemToUpdate).Return(expectedCartItem, nil)

	cartService := NewCartService(cartRepository)

	actualCartItem, err := cartService.Update(ID, itemToUpdate)

	assertion.Nil(err)
	assertion.EqualValues(expectedCartItem, actualCartItem)

	cartRepository.AssertExpectations(t)
}

func TestCartService_Delete(t *testing.T) {
	assertion := assert.New(t)

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("IsCartExists", 123).Return(true, nil)
	cartRepository.On("IsCartItemExists", 123, 456).Return(true, nil)
	cartRepository.On("Delete", 123, 456).Return(nil)

	cartService := CartServiceImpl{rep: cartRepository}

	err := cartService.DeleteItem(123, 456)

	assertion.Nil(err)

	cartRepository.AssertExpectations(t)
}
