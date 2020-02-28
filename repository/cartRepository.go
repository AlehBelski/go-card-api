package repository

import (
	"fmt"
	"github.com/AlehBelski/go-card-api/model"
)

type CartRepository interface {
	Create() (*model.Cart, error)
	Read(id int) (*model.Cart, error)
	Update(id int, item *model.CartItem) (*model.CartItem, error)
	Delete(cartId, itemId int) error
}

func (db *DB) Create() (*model.Cart, error) {
	cart := &model.Cart{}

	query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING id`

	err := db.QueryRow(query).Scan(&cart.Id)
	if err != nil {
		return nil, err
	}
	cart.Items = make([]model.CartItem, 0)

	return cart, nil
}

func (db *DB) Read(id int) (*model.Cart, error) {
	err := db.isCartExists(id)

	if err != nil {
		return nil, err
	}

	query := `
    SELECT * FROM cart_item
    WHERE fk_cart_id = $1`
	rows, err := db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]model.CartItem, 0)
	for rows.Next() {
		var item model.CartItem
		err := rows.Scan(&item.Id, &item.Product, &item.Quantity, &item.Card_id)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	cart := &model.Cart{id, items}

	return cart, nil
}

func (db *DB) Update(id int, item *model.CartItem) (*model.CartItem, error) {
	err := db.isCartExists(id)

	if err != nil {
		return nil, err
	}

	query := `
    INSERT INTO cart_item(product, quantity, fk_cart_id)
    VALUES ($1, $2, $3)
    RETURNING id, fk_cart_id;`

	err = db.QueryRow(query, item.Product, item.Quantity, id).
		Scan(&item.Id, &item.Card_id)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (db *DB) Delete(cartId, itemId int) error {
	err := db.isCartExists(cartId)

	if err != nil {
		return err
	}

	err = db.isCartItemExists(itemId)

	if err != nil {
		return err
	}

	query := `
    DELETE FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

	db.QueryRow(query, cartId, itemId)

	return nil
}

func (db *DB) isCartExists(id int) error {
	query := `
    SELECT 1 FROM cart
    WHERE id = $1`

	var isExist bool
	err := db.QueryRow(query, id).Scan(&isExist)

	if err != nil {
		return fmt.Errorf("specified cart doesn't exists")
	}

	return nil
}

func (db *DB) isCartItemExists(id int) error {
	query := `
    SELECT 1 FROM cart_item
    WHERE id = $1`

	var isExist bool
	err := db.QueryRow(query, id).Scan(&isExist)

	if err != nil {
		return fmt.Errorf("specified cart item doesn't exists")
	}

	return nil
}
