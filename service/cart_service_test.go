package service

import (
    "github.com/AlehBelski/go-card-api/model"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "testing"
)

type CartRepositoryMock struct {
    mock.Mock
}

func (c *CartRepositoryMock) Create() (*model.Cart, error) {
    args := c.Called()

    return args.Get(0).(*model.Cart), args.Error(1)
}

func (c *CartRepositoryMock) Read(id int) (*model.Cart, error) {
    args := c.Called(id)

    return args.Get(0).(*model.Cart), args.Error(1)
}

func (c *CartRepositoryMock) Update(id int, item *model.CartItem) (*model.CartItem, error) {
    args := c.Called(id, item)

    return args.Get(0).(*model.CartItem), args.Error(1)
}

func (c *CartRepositoryMock) Delete(cartId, itemId int) error {
    args := c.Called(cartId, itemId)

    return args.Error(0)
}

func (c *CartRepositoryMock) IsCartExists(id int) (bool, error) {
    args := c.Called(id)

    return args.Bool(0), args.Error(1)
}

func (c *CartRepositoryMock) IsCartItemExists(id int) (bool, error) {
    args := c.Called(id)

    return args.Bool(0), args.Error(1)
}

func TestCartService_Create(t *testing.T) {
    assertion := assert.New(t)

    expectedCart := &model.Cart{
        ID:    123,
        Items: []model.CartItem{},
    }

    cartRepository := new(CartRepositoryMock)

    cartRepository.On("Create").Return(expectedCart, nil)

    cartService := CartService{Rep: cartRepository}

    actualCart, err := cartService.Create()

    assertion.Nil(err)
    assertion.EqualValues(expectedCart, actualCart)

    cartRepository.AssertExpectations(t)
}

func TestCartService_Read(t *testing.T) {
    assertion := assert.New(t)

    expectedCart := &model.Cart{
        ID:    123,
        Items: []model.CartItem{},
    }

    cartRepository := new(CartRepositoryMock)

    cartRepository.On("IsCartExists", expectedCart.ID).Return(true, nil)
    cartRepository.On("Read", expectedCart.ID).Return(expectedCart, nil)

    cartService := CartService{Rep: cartRepository}

    actualCart, err := cartService.Read(expectedCart.ID)

    assertion.Nil(err)
    assertion.EqualValues(expectedCart, actualCart)

    cartRepository.AssertExpectations(t)
}

func TestCartService_Delete(t *testing.T) {
    assertion := assert.New(t)

    cartRepository := new(CartRepositoryMock)

    cartRepository.On("IsCartExists", 123).Return(true, nil)
    cartRepository.On("IsCartItemExists", 456).Return(true, nil)
    cartRepository.On("Delete", 123, 456).Return(nil)

    cartService := CartService{Rep: cartRepository}

    err := cartService.Delete([]string{"/", "cart", "123", "items", "456"})

    assertion.Nil(err)

    cartRepository.AssertExpectations(t)
}
