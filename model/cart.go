package model

type Cart struct {
	Id    int        `json:"id"`
	Items []CartItem `json:"items"`
}
