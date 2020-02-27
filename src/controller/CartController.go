package controller

import (
	"encoding/json"
	"example.com/entry/service"
	"net/http"
	"regexp"
)

var create = regexp.MustCompile("^/carts/?$")
var read = regexp.MustCompile("^/carts/[0-9]+/?$")
var update = regexp.MustCompile("^/carts/[0-9]+/items/?$")
var remove = regexp.MustCompile("^/carts/[0-9]+/items/[0-9]+?$")

func HandleRequest(writer http.ResponseWriter, request *http.Request) {
	switch {
	case create.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleCreate(writer, request)
	case read.MatchString(request.RequestURI) && request.Method == http.MethodGet:
		handleRead(writer, request)
	case update.MatchString(request.RequestURI) && request.Method == http.MethodPost:
		handleUpdate(writer, request)
	case remove.MatchString(request.RequestURI) && request.Method == http.MethodDelete:
		handleRemove(writer, request)
	default:
		http.Error(writer, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

//todo add middleware handler to encode
func handleCreate(writer http.ResponseWriter, request *http.Request) {
	cart := service.Create()

	json.NewEncoder(writer).Encode(cart)
}

func handleRead(writer http.ResponseWriter, request *http.Request) {
	cart := service.Read(request)

	json.NewEncoder(writer).Encode(cart)
}

func handleUpdate(writer http.ResponseWriter, request *http.Request) {
	item := service.Update(request)

	json.NewEncoder(writer).Encode(item)
}

func handleRemove(writer http.ResponseWriter, request *http.Request) {
	service.Delete(request)
}
