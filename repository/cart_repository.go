package repository

import (
	"database/sql"

	"github.com/AlehBelski/go-card-api/model"
)

// Storage contains link to *sql.DB and allow to execute the queries under the database.
type Storage struct {
	db *sql.DB
}

// NewStorage creates new instance of Storage with passed *sql.DB as the argument.
func NewStorage(db *sql.DB) Storage {
	return Storage{db: db}
}

// QueryRow executes a query that is expected to return at most one row.
func (s Storage) QueryRow(query string) {
	s.db.QueryRow(query)
}

// Create creates a new model.Cart record in a database.
// It returns the newly created record as model.Cart.
func (s Storage) Create() (model.Cart, error) {
	cart := model.Cart{}

	query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING ID`

	err := s.db.QueryRow(query).Scan(&cart.ID)
	if err != nil {
		return cart, err
	}
	cart.Items = make([]model.CartItem, 0)

	return cart, nil
}

// Read reads a record from the database and returns it as model.CardDTO.
func (s Storage) Read(ID int) (model.Cart, error) {
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
		err := rows.Scan(&item.ID, &item.Product, &item.Quantity, &item.CartID)
		if err != nil {
			return cart, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return cart, err
	}

	cart.ID = ID
	cart.Items = items

	return cart, nil
}

// Update updates a record in the database related to the passed ID using passed item.
// It returns the updated record as model.CartItem.
func (s Storage) Update(ID int, item model.CartItem) (model.CartItem, error) {
	query := `
    INSERT INTO cart_item(product, quantity, fk_cart_ID)
    VALUES ($1, $2, $3)
    RETURNING ID, fk_cart_ID;`

	err := s.db.QueryRow(query, item.Product, item.Quantity, ID).
		Scan(&item.ID, &item.CartID)

	if err != nil {
		return item, err
	}

	return item, nil
}

// Delete deletes a record from the databases by passed cart and item IDs.
func (s Storage) Delete(cartID, itemID int) error {
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
func (s Storage) IsCartExists(ID int) (bool, error) {
	query := `
    SELECT 1 FROM cart
    WHERE ID = $1`

	var isExist bool
	err := s.db.QueryRow(query, ID).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return true, nil
}

// IsCardExists verify that a model.CartItem object with the specified cartID and itemID presence into a database.
func (s Storage) IsCartItemExists(cartID, itemID int) (bool, error) {
	query := `
    SELECT 1 FROM cart_item
    WHERE fk_cart_ID = $1 AND ID = $2`

	var isExist bool
	err := s.db.QueryRow(query, cartID, itemID).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return true, nil
}
