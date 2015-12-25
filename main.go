package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type mytype []map[string]string

func main() {
	var data mytype
	file, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
