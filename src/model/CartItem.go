package model

type CartItem struct {
	Id       int64  `json:"id"`
	Card_id  int64  `json:"card_id"`
	Product  string `json:"product"`
	Quantity int64  `json:"quantity"`
}

func NewCartItem() CartItem {

	return CartItem{
		Id:       0,
		Card_id:  0,
		Product:  "",
		Quantity: 0,
	}
}
