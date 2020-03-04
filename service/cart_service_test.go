package service

import (
	"testing"

	"github.com/AlehBelski/go-card-api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type CartRepositoryMock struct {
	mock.Mock
}

func (c *CartRepositoryMock) Create() (model.Cart, error) {
	args := c.Called()

	return args.Get(0).(model.Cart), args.Error(1)
}

func (c *CartRepositoryMock) Read(ID int) (model.Cart, error) {
	args := c.Called(ID)

	return args.Get(0).(model.Cart), args.Error(1)
}

func (c *CartRepositoryMock) Update(ID int, item model.CartItem) (model.CartItem, error) {
	args := c.Called(ID, item)

	return args.Get(0).(model.CartItem), args.Error(1)
}

func (c *CartRepositoryMock) Delete(cartID, itemID int) error {
	args := c.Called(cartID, itemID)

	return args.Error(0)
}

func (c *CartRepositoryMock) IsCartExists(ID int) (bool, error) {
	args := c.Called(ID)

	return args.Bool(0), args.Error(1)
}

func (c *CartRepositoryMock) IsCartItemExists(cardID, itemID int) (bool, error) {
	args := c.Called(cardID, itemID)

	return args.Bool(0), args.Error(1)
}

func TestCartService_Create(t *testing.T) {
	assertion := assert.New(t)

	expectedCart := model.Cart{
		ID:    123,
		Items: []model.CartItem{},
	}

	cartRepository := new(CartRepositoryMock)

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

	cartRepository := new(CartRepositoryMock)

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

	cartRepository := new(CartRepositoryMock)

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

	cartRepository := new(CartRepositoryMock)

	cartRepository.On("IsCartExists", 123).Return(true, nil)
	cartRepository.On("IsCartItemExists", 123, 456).Return(true, nil)
	cartRepository.On("Delete", 123, 456).Return(nil)

	cartService := CartServiceImpl{rep: cartRepository}

	err := cartService.DeleteItem(123, 456)

	assertion.Nil(err)

	cartRepository.AssertExpectations(t)
}
