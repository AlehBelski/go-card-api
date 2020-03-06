package controller

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/AlehBelski/go-card-api/model"
)

// Service represent an interface to perform CRUD operations model.Cart object.
type Service interface {
	Create() (model.Cart, error)
	Read(ID int) (model.Cart, error)
	Update(ID int, item model.CartItem) (model.CartItem, error)
	DeleteItem(cartID, itemID int) error
}

// CartController provIDes functionality to handle incoming requests
// for CRUD operations on model.Cart object
type CartController struct {
	service Service
}

// NewCartController creates new CartController object using passed service.CartServiceImpl object.
func NewCartController(service Service) CartController {
	return CartController{service: service}
}

var Create = regexp.MustCompile("^/carts/?$")
var Read = regexp.MustCompile("^/carts/[0-9]+/?$")
var Update = regexp.MustCompile("^/carts/[0-9]+/items/?$")
var Remove = regexp.MustCompile("^/carts/[0-9]+/items/[0-9]+?$")

// HandleCreate handles incoming request to create a new model.CartDTO item and returns it as json string.
func (c CartController) HandleCreate(writer http.ResponseWriter, _ *http.Request) error {
	cart, err := c.service.Create()

	if err != nil {
		return err
	}

	return json.NewEncoder(writer).Encode(cart)
}

// HandleRead handles incoming request to read a model.CartDTO item.
// It retrieves the ID parameters form the request URI and passed it to the next function.
// Returns the result as json string.
func (c CartController) HandleRead(writer http.ResponseWriter, request *http.Request) error {
	ID, err := strconv.Atoi(strings.Split(request.URL.Path, "/")[2])

	if err != nil {
		return err
	}

	cart, err := c.service.Read(ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(writer).Encode(cart)
}

// HandleUpdate handles incoming request to update a model.CartItemDTO item.
// It retrieves the ID parameters form the request URI and passed it to the next function together with request body.
// Returns the result as json string.
func (c CartController) HandleUpdate(writer http.ResponseWriter, request *http.Request) error {
	item := model.CartItem{}
	ID, err := strconv.Atoi(strings.Split(request.URL.Path, "/")[2])

	if err != nil {
		return err
	}

	err = json.NewDecoder(request.Body).Decode(&item)

	if err != nil {
		return err
	}

	item, err = c.service.Update(ID, item)

	if err != nil {
		return err
	}

	return json.NewEncoder(writer).Encode(item)
}

// HandleRemove handle incoming request to remove the specified model.CartItemDTO in the model.CartDTO.
func (c CartController) HandleRemove(_ http.ResponseWriter, request *http.Request) error {
	params := strings.Split(request.URL.Path, "/")
	cartID, err := strconv.Atoi(params[2])

	if err != nil {
		return err
	}

	itemID, err := strconv.Atoi(params[4])

	if err != nil {
		return err
	}

	return c.service.DeleteItem(cartID, itemID)
}
