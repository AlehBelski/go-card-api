package model

// CartItemJson is a copy for model.CartItem object represents the json value.
type CartItemJson struct {
	ID       int    `json:"id"`
	CartID   int    `json:"cart_id"`
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

// NewCartItemJson constructs new CartItemJson object based on passed model.CartItem.
func NewCartItemJson(item CartItem) CartItemJson {
	return CartItemJson{
		ID:       item.ID(),
		CartID:   item.CartID(),
		Product:  item.Product(),
		Quantity: item.Quantity(),
	}
}
