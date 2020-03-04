package model

// CartItem is an object that represents the unique item for shopping cart.
type CartItem struct {
	ID       int
	CartID   int
	Product  string
	Quantity int
}
