package routers

import (
	"net/http"
	productControllers "store/controllers"
)

func LoadRouters() {
	http.HandleFunc("/", productControllers.HandlerIndex)
	http.HandleFunc("/new-product", productControllers.HandlerNewProduct)
	http.HandleFunc("/edit-product", productControllers.HandlerEditProduct)
	http.HandleFunc("/api/product/create", productControllers.HandlerApiCreate)
	http.HandleFunc("/api/product/update", productControllers.HandlerApiUpdate)
	http.HandleFunc("/api/product/delete", productControllers.HandlerApiDelete)
}
