package repository

import (
	"database/sql"
	"example.com/entry/model"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func init() {
	db = OpenAndPingDb()

	query := `
    CREATE TABLE IF NOT EXISTS cart (
        id SERIAL PRIMARY KEY
    );`

	db.QueryRow(query)

	query = `
    CREATE TABLE IF NOT EXISTS cart_item (
        id SERIAL PRIMARY KEY,
        product TEXT NOT NULL,
        quantity INT NOT NULL,
        fk_cart_id INT REFERENCES cart(id)
    );`

	db.QueryRow(query)
}

func Create() model.Cart {
	query := `
    INSERT INTO cart
    VALUES(DEFAULT)
    RETURNING id`

	cart := model.NewCart()
	db.QueryRow(query).Scan(&cart.Id)
	cart.Items = make([]model.CartItem, 0) //todo ?

	return cart
}

//fixme add exception handlers
func Read(id int) model.Cart {
	items := make([]model.CartItem, 0)

	query := `
    SELECT * FROM cart_item
    WHERE fk_cart_id = $1`
	rows, _ := db.Query(query, id)

	for rows.Next() {
		item := model.NewCartItem()
		err := rows.Scan(&item.Id, &item.Product, &item.Quantity, &item.Card_id)
		if err != nil {
			log.Fatal(err)
		}

		items = append(items, item)
	}

	cart := model.NewCart()
	cart.Id = id
	cart.Items = items

	return cart
}

func Update(id int, item model.CartItem) model.CartItem {
	query := `
    INSERT INTO cart_item(product, quantity, fk_cart_id)
    VALUES ($1, $2, $3)
    RETURNING id, fk_cart_id;`

	db.QueryRow(query, item.Product, item.Quantity, id).
		Scan(&item.Id, &item.Card_id)

	return item
}

func Delete(cartId, itemId int) {
	query := `
    DELETE FROM cart_item
    WHERE fk_cart_id = $1 AND id = $2`

	db.QueryRow(query, cartId, itemId)

	return

}
