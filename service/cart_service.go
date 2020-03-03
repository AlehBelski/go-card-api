package service

import (
    "fmt"
    "github.com/AlehBelski/go-card-api/model"
    "strings"
)

// CartRepository represent an interface to perform CRUD operations under model.Cart object.
// Moreover, it has additional functions to verify model.Cart and model.CartItem presence.
type CartRepository interface {
    Create() (*model.Cart, error)
    Read(id int) (*model.Cart, error)
    Update(id int, item model.CartItem) (model.CartItem, error)
    Delete(cartId, itemId int) error
    IsCartExists(id int) (bool, error)
    IsCartItemExists(cartId, itemId int) (bool, error)
}

// CartService used to represent a business logic for model.Cart object.
type CartService struct {
    rep CartRepository
}

// NewCartService creates new CartService object using passed repository.CartRepository object.
func NewCartService(cartRepository CartRepository) CartService {
    return CartService{cartRepository}
}

// Create calls srv.rep.Create() in order to create a new record in a database.
func (srv CartService) Create() (*model.Cart, error) {
    return srv.rep.Create()
}

// Read calls srv.rep.Read() in order to read a record from a database.
func (srv CartService) Read(id int) (*model.Cart, error) {
    exists, err := srv.rep.IsCartExists(id)

    if err != nil {
        return nil, err
    }

    if !exists {
        return nil, fmt.Errorf("specified cart %v doesn't exists", id)
    }

    return srv.rep.Read(id)
}

// Update calls srv.rep.Update() in order to update a record in database by passed id using the request body.
// It returns an error in case of item.Product is blank or item.Quantity is not positive.
func (srv CartService) Update(id int, item model.CartItem) (model.CartItem, error) {
    exists, err := srv.rep.IsCartExists(id)

    if err != nil || !exists {
        return item, fmt.Errorf("specified cart %v doesn't exists", id)
    }

    if len(strings.TrimSpace(item.Product)) == 0 {
        return item, fmt.Errorf("product: %q name should not be blank", item.Product)
    }

    if item.Quantity <= 0 {
        return item, fmt.Errorf("quantity: %q should be a positive value", item.Quantity)
    }

    return srv.rep.Update(id, item)
}

// Delete converts the passed params in order to retrieve cart and item id
// and calls srv.rep.Delete to delete the specified record.
func (srv CartService) Delete(cartId, itemId int) error {
    exists, err := srv.rep.IsCartExists(cartId)

    if err != nil || !exists {
        return fmt.Errorf("specified cart %v doesn't exists", cartId)
    }

    exists, err = srv.rep.IsCartItemExists(cartId, itemId)

    if err != nil || !exists {
        return fmt.Errorf("specified cart item %v doesn't exists", cartId)
    }

    return srv.rep.Delete(cartId, itemId)
}
