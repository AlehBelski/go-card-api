package service

import (
	"fmt"
	"strings"

	"github.com/AlehBelski/go-card-api/model"
)

// CartRepository represent an interface to perform CRUD operations under database for model.Cart object.
// Moreover, it has additional functions to verify model.Cart and model.CartItem presence.
type CartRepository interface {
	Create() (model.Cart, error)
	Read(ID int) (model.Cart, error)
	Update(ID int, item model.CartItem) (model.CartItem, error)
	Delete(cartID, itemID int) error
	IsCartExists(ID int) (bool, error)
	IsCartItemExists(cartID, itemID int) (bool, error)
}

// CartServiceImpl used to represent a business logic for model.Cart object.
type CartServiceImpl struct {
	rep CartRepository
}

// NewCartService creates new CartServiceImpl object using passed repository.CartRepository object.
func NewCartService(cartRepository CartRepository) CartServiceImpl {
	return CartServiceImpl{cartRepository}
}

// Create calls srv.rep.Create() in order to create a new record in a database.
func (srv CartServiceImpl) Create() (model.Cart, error) {
	return srv.rep.Create()
}

// Read calls srv.rep.Read() in order to read a record from a database.
func (srv CartServiceImpl) Read(ID int) (model.Cart, error) {
	cart := model.Cart{}
	exists, err := srv.rep.IsCartExists(ID)

	if err != nil {
		return cart, err
	}

	if !exists {
		return cart, fmt.Errorf("specified cart %v doesn't exists", ID)
	}

	return srv.rep.Read(ID)
}

// Update calls srv.rep.Update() in order to update a record in database by passed ID using the request body.
// It returns an error in case of item.Product is blank or item.Quantity is not positive.
func (srv CartServiceImpl) Update(ID int, item model.CartItem) (model.CartItem, error) {
	exists, err := srv.rep.IsCartExists(ID)

	if err != nil || !exists {
		return item, fmt.Errorf("specified cart %v doesn't exists", ID)
	}

	if len(strings.TrimSpace(item.Product)) == 0 {
		return item, fmt.Errorf("product: %q name should not be blank", item.Product)
	}

	if item.Quantity <= 0 {
		return item, fmt.Errorf("quantity: %q should be a positive value", item.Quantity)
	}

	return srv.rep.Update(ID, item)
}

// DeleteItem converts the passed params in order to retrieve cart and item ID
// and calls srv.rep.Delete to delete the specified record.
func (srv CartServiceImpl) DeleteItem(cardID, itemID int) error {
	exists, err := srv.rep.IsCartExists(cardID)

	if err != nil || !exists {
		return fmt.Errorf("specified cart %v doesn't exists", cardID)
	}

	exists, err = srv.rep.IsCartItemExists(cardID, itemID)

	if err != nil || !exists {
		return fmt.Errorf("specified cart item %v doesn't exists", cardID)
	}

	return srv.rep.Delete(cardID, itemID)
}
