package main

import (
	"fmt"
	"os"

	"github.com/gofor-little/env"
)

func main() {
	fmt.Println("Working with Environments variables")
	fmt.Println("GOPATH:", os.Getenv("GOPATH"))

	err := env.Load("environments/.env")
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		databasePass := env.Get("DATABASE_PASS", "")
		fmt.Println("DATABASE_PASS:", databasePass)
	}
}
