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
	db, err := newDB("postgres", "postgres", "pstgr", "postgres")
	if err != nil {
		panic(err)
	}

	env := Env{controller.NewCartController(service.NewCartService(db))}

	http.HandleFunc("/", middleware.LogMiddleware(env.handleRequest))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newDB(userName, userPassword, host, dbName string) (repository.CartRepositoryImpl, error) {
	storage := repository.CartRepositoryImpl{}
	dataSource := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", userName, userPassword, host, dbName)
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return storage, err
	}
	if err = db.Ping(); err != nil {
		return storage, err
	}
	storage = repository.NewStorage(db)
	return storage, nil
}

func (env Env) handleRequest(writer http.ResponseWriter, request *http.Request) {
	switch {
	case controller.CreateURL.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleOperation(env.controller.HandleCreate, writer, request)
	case controller.ReadURL.MatchString(request.RequestURI) && request.Method == http.MethodGet:
		handleOperation(env.controller.HandleRead, writer, request)
	case controller.UpdateURL.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleOperation(env.controller.HandleUpdate, writer, request)
	case controller.RemoveURL.MatchString(request.RequestURI) && request.Method == http.MethodDelete:
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
