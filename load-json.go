package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// slice of maps. each map hase k/v that both are strings
type mytype []map[string]string

func LoadJSON() mytype {
	var data mytype
	file, err := ioutil.ReadFile("products-small.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
