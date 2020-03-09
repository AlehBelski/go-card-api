// Package model contains base objects over which the main implementation is working on.
package model

// Cart is an object that represents an online shopping cart.
type Cart struct {
	id    int
	items []CartItem
}

// NewCart creates new Cart object based on the passed ID and items.
func NewCart(id int, items []CartItem) Cart {
	return Cart{
		id:    id,
		items: items,
	}
}

// SetID sets id of the Cart object.
func (c *Cart) SetID(id int) {
	c.id = id
}

// ID returns the id of the Cart object.
func (c Cart) ID() int {
	return c.id
}

// SetItems sets items of the Cart object.
func (c *Cart) SetItems(items []CartItem) {
	c.items = items
}

// Items returns the items of the Cart object.
func (c Cart) Items() []CartItem {
	return c.items
}
