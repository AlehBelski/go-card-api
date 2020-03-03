package repository

import (
    "database/sql"
    "github.com/AlehBelski/go-card-api/model"
)

type DB struct {
    *sql.DB
}

// Create creates a new model.Cart record in a database.
// It returns the newly created record as *model.Cart.
func (db *DB) Create() (*model.Cart, error) {
    cart := &model.Cart{}

    query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING id`

    err := db.QueryRow(query).Scan(&cart.ID)
    if err != nil {
        return nil, err
    }
    cart.Items = make([]model.CartItem, 0)

    return cart, nil
}

// Read reads a record from the database and returns it as *model.CardDTO.
func (db *DB) Read(id int) (*model.Cart, error) {
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
// It returns the updated record as *model.CartItem.
func (db *DB) Update(id int, item model.CartItem) (model.CartItem, error) {
    cartItem := model.CartItem{}

    query := `
    INSERT INTO cart_item(product, quantity, fk_cart_id)
    VALUES ($1, $2, $3)
    RETURNING id, fk_cart_id;`

    err := db.QueryRow(query, item.Product, item.Quantity, id).
        Scan(&cartItem.ID, &cartItem.CardId)

    if err != nil {
        return cartItem, err
    }

    return cartItem, nil
}

// Delete deletes a record from the databases by passed cart and item ids.
func (db *DB) Delete(cartId, itemId int) error {
    query := `
    DELETE FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

    db.QueryRow(query, cartId, itemId)

    return nil
}

// IsCardExists verify that a model.Cart object with the specified id presence into a database.
func (db *DB) IsCartExists(id int) (bool, error) {
    query := `
    SELECT 1 FROM cart
    WHERE id = $1`

    var isExist bool
    err := db.QueryRow(query, id).Scan(&isExist)

    if err != nil {
        return false, err
    }

    return true, nil
}

// IsCardExists verify that a model.CartItem object with the specified cartId and itemId presence into a database.
func (db *DB) IsCartItemExists(cartId, itemId int) (bool, error) {
    query := `
    SELECT 1 FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

    var isExist bool
    err := db.QueryRow(query, cartId, itemId).Scan(&isExist)

    if err != nil {
        return false, err
    }

    return true, nil
}