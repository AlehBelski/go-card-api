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

// Create calls srv.Rep.Create() in order to create a new record in a database.
func (srv CartService) Create() (*model.CartDTO, error) {
    return srv.Rep.Create()
}

// Read calls srv.Rep.Read() in order to read a record from a database.
func (srv CartService) Read(id int) (*model.CartDTO, error) {
    return srv.Rep.Read(id)
}

// Update calls srv.Rep.Update() in order to update a record in database by passed id using the request body.
// It returns an error in case of item.Product is blank or item.Quantity is not positive.
func (srv CartService) Update(id int, body io.ReadCloser) (*model.CartItemDTO, error) {
    item := &model.CartItemDTO{}

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

// Delete converts the passed params in order to retrieve cart and item id
// and calls srv.Rep.Delete to delete the specified record.
func (srv CartService) Delete(params []string) error {
    cartId, _ := strconv.Atoi(params[2])
    itemId, _ := strconv.Atoi(params[4])
    return srv.Rep.Delete(cartId, itemId)
}
