package model

// CartJson is a copy for model.Cart object represents the json value.
type CartJson struct {
	ID    int            `json:"id"`
	Items []CartItemJson `json:"items"`
}

// NewCartJson constructs new CartItemJson object based on passed model.Cart.
func NewCartJson(item Cart) CartJson {
	items := make([]CartItemJson, len(item.Items()))

	for i, v := range item.Items() {
		items[i] = NewCartItemJson(v)
	}

	return CartJson{
		ID:    item.ID(),
		Items: items,
	}
}
