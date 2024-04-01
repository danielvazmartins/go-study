package productControllers

import (
	"fmt"
	"html/template"
	"net/http"
	productModel "store/models"
	"strconv"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	products := productModel.GetAll()
	templates.ExecuteTemplate(w, "Index", products)
}

func HandlerNewProduct(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "NewProduct", nil)
}

func HandlerEditProduct(w http.ResponseWriter, r *http.Request) {
	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		fmt.Println("Value of id is invalid")
	}

	product := productModel.GetProductById(id)
	templates.ExecuteTemplate(w, "EditProduct", product)
}

func HandlerApiCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceRaw := r.FormValue("price")
		amountRaw := r.FormValue("amount")

		price, err := strconv.ParseFloat(priceRaw, 64)
		if err != nil {
			fmt.Println("Value of price is invalid")
		}

		amount, err := strconv.Atoi(amountRaw)
		if err != nil {
			fmt.Println("Value of amount is invalid")
		}
		productModel.Create(name, description, price, amount)

		http.Redirect(w, r, "/", 302)
	}
}

func HandlerApiUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		idRaw := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		priceRaw := r.FormValue("price")
		amountRaw := r.FormValue("amount")

		id, err := strconv.Atoi(idRaw)
		if err != nil {
			fmt.Println("Value of id is invalid")
		}

		price, err := strconv.ParseFloat(priceRaw, 64)
		if err != nil {
			fmt.Println("Value of price is invalid")
		}

		amount, err := strconv.Atoi(amountRaw)
		if err != nil {
			fmt.Println("Value of amount is invalid")
		}
		productModel.Update(id, name, description, price, amount)

		http.Redirect(w, r, "/", 302)
	}
}

func HandlerApiDelete(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.URL.Query().Get("id"))
	idRaw := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		fmt.Println("Value of id is invalid")
	}
	productModel.Delete(id)

	http.Redirect(w, r, "/", 302)
}
