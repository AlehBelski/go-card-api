package service

import (
	"encoding/json"
	"example.com/entry/model"
	"example.com/entry/repository"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//fixme add exception handlers
func Create() model.Cart {
	return repository.Create()
}

func Read(request *http.Request) model.Cart {
	id, _ := strconv.Atoi(strings.Split(request.RequestURI, "/")[2])

	return repository.Read(id)
}

func Update(request *http.Request) model.CartItem {
	var item model.CartItem

	b, _ := ioutil.ReadAll(request.Body)
	json.Unmarshal(b, &item)

	id, _ := strconv.Atoi(strings.Split(request.RequestURI, "/")[2])

	return repository.Update(id, item)
}

func Delete(request *http.Request) {
	params := strings.Split(request.RequestURI, "/")
	cartId, _ := strconv.Atoi(params[2])
	itemId, _ := strconv.Atoi(params[4])
	repository.Delete(cartId, itemId)
}
