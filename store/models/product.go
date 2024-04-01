package productModel

import (
	"fmt"
	"store/db"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func GetAll() []Product {
	fmt.Println("Getting products...")

	db := db.Init()
	dbResult, err := db.Query("select * from products order by id")
	if err != nil {
		fmt.Println("Error:", err)
	}

	var products []Product

	for dbResult.Next() {
		var id int
		var name string
		var description string
		var price float64
		var amount int

		err := dbResult.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err)
		}

		product := Product{
			Id:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Amount:      amount,
		}
		products = append(products, product)
	}
	return products
}
func GetProductById(id int) Product {
	fmt.Println("Getting product...")

	db := db.Init()
	dbResult, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		fmt.Println("Error:", err)
	}

	var product Product
	for dbResult.Next() {
		err = dbResult.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Amount)
		if err != nil {
			panic(err)
		}
	}

	return product
}

func Create(name string, description string, price float64, amount int) {
	fmt.Println("Creating product...")

	db := db.Init()
	dbResult, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		dbResult.Exec(name, description, price, amount)
	}
}

func Update(id int, name string, description string, price float64, amount int) {
	fmt.Println("Updating product...")

	db := db.Init()
	dbResult, err := db.Prepare("update products set name=$2, description=$3, price=$4, amount=$5 where id=$1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		dbResult.Exec(id, name, description, price, amount)
	}

}

func Delete(id int) {
	fmt.Println("Deleting product...")

	db := db.Init()
	dbResult, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		dbResult.Exec(id)
	}
}
