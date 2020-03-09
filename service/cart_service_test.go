package service

import (
	"errors"
	"testing"

	"github.com/AlehBelski/go-card-api/model"
	"github.com/AlehBelski/go-card-api/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestCartService_Create(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.NewCart(123, []model.CartItem{})

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("Create").Return(expectedCart, nil)

	cartService := NewCartService(cartRepository)

	actualCart, err := cartService.Create()

	assertion.Nil(err)
	assertion.EqualValues(expectedCart, actualCart)

	cartRepository.AssertExpectations(t)
}

var readTests = []struct {
	ID       int
	isExists bool
	out      model.Cart
	err      error
}{
	{123, true, model.NewCart(123, []model.CartItem{}), nil},
	{444, false, model.Cart{}, errors.New("specified cart 444 doesn't exists")},
}

func TestCartService_Read(t *testing.T) {
	assertion := assert.New(t)

	for _, v := range readTests {
		t.Run("test read function", func(t *testing.T) {
			cartRepository := new(mocks.CartRepository)

			cartRepository.On("IsCartExists", v.ID).Return(v.isExists, nil)

			if v.isExists {
				cartRepository.On("Read", v.ID).Return(v.out, v.err)
			}

			cartService := NewCartService(cartRepository)

			actualCart, err := cartService.Read(v.ID)

			assertion.Equal(v.err, err)
			assertion.EqualValues(v.out, actualCart)

			cartRepository.AssertExpectations(t)
		})
	}
}

func TestCartService_Update(t *testing.T) {
	assertion := assert.New(t)

	ID := 123

	itemToUpdate := model.CartItem{}

	itemToUpdate.SetProduct("Shoes")
	itemToUpdate.SetQuantity(100)

	expectedCartItem := model.CartItem{}

	expectedCartItem.SetID(ID)
	expectedCartItem.SetProduct(itemToUpdate.Product())
	expectedCartItem.SetQuantity(itemToUpdate.Quantity())

	cartRepository := new(mocks.CartRepository)

	cartRepository.On("IsCartExists", ID).Return(true, nil)
	cartRepository.On("Update", ID, itemToUpdate).Return(expectedCartItem, nil)

	cartService := NewCartService(cartRepository)

	actualCartItem, err := cartService.Update(ID, itemToUpdate)

	assertion.Nil(err)
	assertion.EqualValues(expectedCartItem, actualCartItem)

	cartRepository.AssertExpectations(t)
}

var deleteTests = []struct {
	cartID           int
	isCartExists     bool
	itemID           int
	isCartItemExists bool
	err              error
}{
	{123, true, 456, true, nil},
	{123, true, 444, false, errors.New("specified cart item 444 doesn't exists")},
	{444, false, 444, true, errors.New("specified cart 444 doesn't exists")},
}

func TestCartService_Delete(t *testing.T) {
	assertion := assert.New(t)

	for _, v := range deleteTests {
		t.Run("test delete function", func(t *testing.T) {
			cartRepository := new(mocks.CartRepository)

			cartRepository.On("IsCartExists", v.cartID).Return(v.isCartExists, nil)

			if v.isCartExists {
				cartRepository.On("IsCartItemExists", v.cartID, v.itemID).Return(v.isCartItemExists, nil)
			}

			if v.isCartExists && v.isCartItemExists {
				cartRepository.On("Delete", v.cartID, v.itemID).Return(v.err)
			}

			cartService := CartServiceImpl{rep: cartRepository}

			err := cartService.DeleteItem(v.cartID, v.itemID)

			assertion.Equal(v.err, err)

			cartRepository.AssertExpectations(t)
		})
	}

}
