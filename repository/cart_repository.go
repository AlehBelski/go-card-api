package repository

import (
	"database/sql"

	"github.com/AlehBelski/go-card-api/model"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

func (s Storage) DB() *sql.DB {
	return s.db
}

// Create creates a new model.Cart record in a database.
// It returns the newly created record as *model.Cart.
func (s *Storage) Create() (*model.Cart, error) {
	cart := &model.Cart{}

	query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING id`

	err := s.db.QueryRow(query).Scan(&cart.ID)
	if err != nil {
		return nil, err
	}
	cart.Items = make([]model.CartItem, 0)

	return cart, nil
}

// Read reads a record from the database and returns it as *model.CardDTO.
func (s *Storage) Read(id int) (*model.Cart, error) {
	query := `
    SELECT * FROM cart_item
    WHERE fk_cart_id = $1`
	rows, err := s.db.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := make([]model.CartItem, 0)
	for rows.Next() {
		var item model.CartItem
		err := rows.Scan(&item.ID, &item.Product, &item.Quantity, &item.CardId)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	cart := &model.Cart{ID: id, Items: items}

	return cart, nil
}

// Update updates a record in the database related to the passed id using passed item.
// It returns the updated record as model.CartItem.
func (s *Storage) Update(id int, item model.CartItem) (model.CartItem, error) {
	query := `
    INSERT INTO cart_item(product, quantity, fk_cart_id)
    VALUES ($1, $2, $3)
    RETURNING id, fk_cart_id;`

	err := s.db.QueryRow(query, item.Product, item.Quantity, id).
		Scan(&item.ID, &item.CardId)

	if err != nil {
		return item, err
	}

	return item, nil
}

// Delete deletes a record from the databases by passed cart and item ids.
func (s *Storage) Delete(cartId, itemId int) error {
	query := `
    DELETE FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

	rows, err := s.db.Query(query, cartId, itemId)

	if err != nil {
		return err
	}

	if err = rows.Close(); err != nil {
		return err
	}

	return nil
}

// IsCardExists verify that a model.Cart object with the specified id presence into a database.
func (s *Storage) IsCartExists(id int) (bool, error) {
	query := `
    SELECT 1 FROM cart
    WHERE id = $1`

	var isExist bool
	err := s.db.QueryRow(query, id).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return true, nil
}

// IsCardExists verify that a model.CartItem object with the specified cartId and itemId presence into a database.
func (s *Storage) IsCartItemExists(cartId, itemId int) (bool, error) {
	query := `
    SELECT 1 FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

	var isExist bool
	err := s.db.QueryRow(query, cartId, itemId).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return true, nil
}
