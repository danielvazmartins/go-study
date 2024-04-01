package main

import (
	"fmt"
	"net/http"
	product "store/models"
	"store/routers"
)

var products []product.Product

func main() {
	routers.LoadRouters()

	addrPort := ":8080"
	fmt.Println("Starting store in http://localhost" + addrPort)
	http.ListenAndServe(addrPort, nil)
}
