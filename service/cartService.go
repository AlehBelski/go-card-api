package service

import (
	"encoding/json"
	"fmt"
	"github.com/AlehBelski/go-card-api/model"
	"github.com/AlehBelski/go-card-api/repository"
	"io"
	"strconv"
)

type CartService struct {
	Rep repository.CartRepository
}

func (srv CartService) Create() (*model.Cart, error) {
	return srv.Rep.Create()
}

func (srv CartService) Read(id int) (*model.Cart, error) {
	return srv.Rep.Read(id)
}

func (srv CartService) Update(id int, body io.ReadCloser) (*model.CartItem, error) {
	item := &model.CartItem{}

	err := json.NewDecoder(body).Decode(&item)

	if err != nil {
		return nil, err
	}

	if len(item.Product) == 0 {
		return nil, fmt.Errorf("product name should not be blank")
	}

	if item.Quantity < 0 {
		return nil, fmt.Errorf("quantity should be a positive value")
	}

	return srv.Rep.Update(id, item)
}

func (srv CartService) Delete(params []string) error {
	cartId, _ := strconv.Atoi(params[2])
	itemId, _ := strconv.Atoi(params[4])
	return srv.Rep.Delete(cartId, itemId)
}
