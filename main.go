package main

import (
    "github.com/AlehBelski/go-card-api/controller"
    "github.com/AlehBelski/go-card-api/handler"
    "github.com/AlehBelski/go-card-api/repository"
    "github.com/AlehBelski/go-card-api/service"
    _ "github.com/lib/pq"
    "log"
    "net/http"
)

type Env struct {
    controller *controller.CartController
}

func main() {
    db, err := repository.NewDB("postgres", "postgres", "localhost", "postgres")
    if err != nil {
        panic(err)
    }
    initDb(db)

    env := &Env{&controller.CartController{Service: &service.CartService{Rep: db}}}

    http.HandleFunc("/", handler.LogHandler(env.handleRequest))

    err = http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatalf("not able to start the server: %s", err)
    }
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
