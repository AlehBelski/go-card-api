package model

// CartItem is an object that represents the unique item for shopping cart.
type CartItem struct {
	id       int
	cartID   int
	product  string
	quantity int
}

// NewCartItem creates new CartItem object based on the passed values.
func NewCartItem(id int, cartID int, product string, quantity int) CartItem {
	return CartItem{
		id:       id,
		cartID:   cartID,
		product:  product,
		quantity: quantity,
	}
}

// SetID sets ID of the CartItem Object using the passed value.
func (c *CartItem) SetID(id int) {
	c.id = id
}

// ID return the value of ID field of the CartItem object.
func (c *CartItem) ID() int {
	return c.id
}

// SetCartID sets cartID of the CartItem Object using the passed value.
func (c *CartItem) SetCartID(cartID int) {
	c.cartID = cartID
}

// CartID return the value of cartID field of the CartItem object.
func (c *CartItem) CartID() int {
	return c.cartID
}

// SetProduct sets product of the CartItem Object using the passed value.
func (c *CartItem) SetProduct(product string) {
	c.product = product
}

// Product return the value of product field of the CartItem object.
func (c *CartItem) Product() string {
	return c.product
}

// SetQuantity sets quantity of the CartItem Object using the passed value.
func (c *CartItem) SetQuantity(quantity int) {
	c.quantity = quantity
}

// Quantity return the value of quantity field of the CartItem object.
func (c *CartItem) Quantity() int {
	return c.quantity
}

// NewCartItemFromJson constructs new CartItem object based on the CartItemJson value.
func NewCartItemFromJson(json CartItemJson) CartItem {
	return CartItem{
		id:       json.ID,
		cartID:   json.CartID,
		product:  json.Product,
		quantity: json.Quantity,
	}
}
