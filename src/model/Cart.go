package model

type Cart struct {
	Id    int        `json:"id"`
	Items []CartItem `json:"items"`
}

func NewCart() Cart {

	return Cart{
		Id:    0,
		Items: []CartItem{},
	}
}
