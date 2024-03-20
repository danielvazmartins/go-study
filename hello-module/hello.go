package main

import (
	"fmt"

	"example/hello/sub"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Go with modules")
	quoteGo := quote.Go()
	fmt.Println("quote Go:", quoteGo)

	version := sub.SubVersion()
	fmt.Println("Versão:", version)
}
