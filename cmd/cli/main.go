package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/oren/search"
)

var Products *search.Products

func main() {
	var err error
	Products, err = search.New("products.xml")
	if err != nil {
		panic(err)
	}
	query := strings.Join(os.Args, " ")
	if query != "" {
		results := Products.Search(query)
		data, err := json.Marshal(results)
		if err != nil {
			fmt.Errorf("encode response: %v", err)
		}
		os.Stdout.Write(data)
	}
}
