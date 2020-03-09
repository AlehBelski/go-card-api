// Package repository represents operations over a database.
package repository

import (
	"database/sql"

	"github.com/AlehBelski/go-card-api/model"
)

// CartRepositoryImpl contains link to *sql.DB and allow to execute the queries under the database.
type CartRepositoryImpl struct {
	db *sql.DB
}

// NewStorage creates new instance of CartRepositoryImpl with passed *sql.DB as the argument.
func NewStorage(db *sql.DB) CartRepositoryImpl {
	return CartRepositoryImpl{db: db}
}

// QueryRow executes a query that is expected to return at most one row.
func (s CartRepositoryImpl) QueryRow(query string) {
	s.db.QueryRow(query)
}

// Create creates a new model.Cart record in a database.
// It returns the newly created record as model.Cart.
func (s CartRepositoryImpl) Create() (model.Cart, error) {
	cart := model.Cart{}

	query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING ID`

	var id int
	err := s.db.QueryRow(query).Scan(&id)
	if err != nil {
		return cart, err
	}

	cart.SetID(id)
	cart.SetItems(make([]model.CartItem, 0))

	return cart, nil
}

// Read reads a record from the database and returns it as model.CardDTO.
func (s CartRepositoryImpl) Read(ID int) (model.Cart, error) {
	cart := model.Cart{}

	query := `
    SELECT * FROM cart_item
    WHERE fk_cart_ID = $1`
	rows, err := s.db.Query(query, ID)

	if err != nil {
		return cart, err
	}

	defer rows.Close()

	items := make([]model.CartItem, 0)
	for rows.Next() {
		var item model.CartItem

		var id int
		var cartID int
		var product string
		var quantity int

		err := rows.Scan(&id, &product, &quantity, &cartID)
		if err != nil {
			return cart, err
		}

		item.SetID(id)
		item.SetCartID(cartID)
		item.SetProduct(product)
		item.SetQuantity(quantity)

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return cart, err
	}

	cart.SetID(ID)
	cart.SetItems(items)

	return cart, nil
}

// Update updates a record in the database related to the passed ID using passed item.
// It returns the updated record as model.CartItem.
func (s CartRepositoryImpl) Update(ID int, item model.CartItem) (model.CartItem, error) {
	query := `
    INSERT INTO cart_item(product, quantity, fk_cart_ID)
    VALUES ($1, $2, $3)
    RETURNING ID, fk_cart_ID;`

	var id int
	var cartID int

	err := s.db.QueryRow(query, item.Product(), item.Quantity(), ID).
		Scan(&id, &cartID)

	if err != nil {
		return item, err
	}

	item.SetID(id)
	item.SetCartID(cartID)

	return item, nil
}

// Delete deletes a record from the databases by passed cart and item IDs.
func (s CartRepositoryImpl) Delete(cartID, itemID int) error {
	query := `
    DELETE FROM cart_item
    WHERE fk_cart_ID = $1 AND ID = $2`

	rows, err := s.db.Query(query, cartID, itemID)

	if err != nil {
		return err
	}

	return rows.Close()
}

// IsCardExists verify that a model.Cart object with the specified ID presence into a database.
func (s CartRepositoryImpl) IsCartExists(ID int) (bool, error) {
	query := `
    SELECT 1 FROM cart
    WHERE ID = $1`

	rows, err := s.db.Query(query, ID)

	if err != nil {
		return false, err
	}

	return rows.Next(), err
}

// IsCardExists verify that a model.CartItem object with the specified cartID and itemID presence into a database.
func (s CartRepositoryImpl) IsCartItemExists(cartID, itemID int) (bool, error) {
	query := `
    SELECT 1 FROM cart_item
    WHERE fk_cart_ID = $1 AND ID = $2`

	rows, err := s.db.Query(query, cartID, itemID)

	if err != nil {
		return false, err
	}

	return rows.Next(), err
}
