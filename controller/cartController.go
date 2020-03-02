package controller

import (
    "encoding/json"
    "github.com/AlehBelski/go-card-api/service"
    "net/http"
    "regexp"
    "strconv"
    "strings"
)

type CartController struct {
    Service *service.CartService
}

var Create = regexp.MustCompile("^/carts/?$")
var Read = regexp.MustCompile("^/carts/[0-9]+/?$")
var Update = regexp.MustCompile("^/carts/[0-9]+/items/?$")
var Remove = regexp.MustCompile("^/carts/[0-9]+/items/[0-9]+?$")

// HandleCreate handles incoming request to create a new model.CartDTO item and returns it as json string.
func (c CartController) HandleCreate(writer http.ResponseWriter, request *http.Request) error {
    cart, err := c.Service.Create()

    if err != nil {
        return err
    }

    err = json.NewEncoder(writer).Encode(cart)

    return err
}

// HandleRead handles incoming request to read a model.CartDTO item.
// It retrieves the id parameters form the request URI and passed it to the next function.
// Returns the result as json string.
func (c CartController) HandleRead(writer http.ResponseWriter, request *http.Request) error {
    id, _ := strconv.Atoi(strings.Split(request.RequestURI, "/")[2])
    cart, err := c.Service.Read(id)

    if err != nil {
        return err
    }

    err = json.NewEncoder(writer).Encode(cart)

    return err
}

// HandleUpdate handles incoming request to update a model.CartItemDTO item.
// It retrieves the id parameters form the request URI and passed it to the next function together with request body.
// Returns the result as json string.
func (c CartController) HandleUpdate(writer http.ResponseWriter, request *http.Request) error {
    id, _ := strconv.Atoi(strings.Split(request.RequestURI, "/")[2])
    item, err := c.Service.Update(id, request.Body)

    if err != nil {
        return err
    }

    err = json.NewEncoder(writer).Encode(item)

    return err
}

// HandleRemove handle incoming request to remove the specified model.CartItemDTO in the model.CartDTO.
func (c CartController) HandleRemove(writer http.ResponseWriter, request *http.Request) error {
    params := strings.Split(request.RequestURI, "/")

    return c.Service.Delete(params)
}
