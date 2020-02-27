package main

import (
	"example.com/entry/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/", controller.HandleRequest)

	http.ListenAndServe(":3000", nil)
}
