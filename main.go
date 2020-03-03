package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/AlehBelski/go-card-api/controller"
	"github.com/AlehBelski/go-card-api/middleware"
	"github.com/AlehBelski/go-card-api/repository"
	"github.com/AlehBelski/go-card-api/service"
	_ "github.com/lib/pq"
)

type Env struct {
	controller controller.CartController
}

func main() {
	db, err := newDB("postgres", "postgres", "localhost", "postgres")
	if err != nil {
		panic(err)
	}
	initDb(db)

	env := &Env{controller.NewCartController(service.NewCartService(db))}

	http.HandleFunc("/", middleware.LogMiddleware(env.handleRequest))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newDB(userName, userPassword, host, dbName string) (*repository.DB, error) {
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", userName, userPassword, host, dbName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &repository.DB{DB: db}, nil
}

func initDb(db *repository.DB) {
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

func (env Env) handleRequest(writer http.ResponseWriter, request *http.Request) {
	switch {
	case controller.Create.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleOperation(env.controller.HandleCreate, writer, request)
	case controller.Read.MatchString(request.RequestURI) && request.Method == http.MethodGet:
		handleOperation(env.controller.HandleRead, writer, request)
	case controller.Update.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleOperation(env.controller.HandleUpdate, writer, request)
	case controller.Remove.MatchString(request.RequestURI) && request.Method == http.MethodDelete:
		handleOperation(env.controller.HandleRemove, writer, request)
	default:
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func handleOperation(fn func(w http.ResponseWriter, r *http.Request) error, w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}
