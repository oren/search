package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
)

type DataFormat struct {
	Rss struct {
		Channel struct {
			ProductList []struct {
				Sku      string `xml:"sku" json:"sku"`
				Quantity int    `xml:"quantity" json:"quantity"`
			} `xml:"Product" json:"products"`
		} `xml:"ProductList" json:"ProductList"`
	} `xml:"channel" json:"channel"`
}

func main() {
	xmlData, err := ioutil.ReadFile("products.xml")
	if err != nil {
		log.Fatal(err)
	}
	data := &DataFormat{}

	err = xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}

	fmt.Printf("%s\n", data)

	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	fmt.Printf("%s\n", result)
}
