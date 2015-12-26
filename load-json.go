package main

import (
	// "encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type DataFormat struct {
	Rss struct {
		Channel []struct {
			Id          string `xml:"id" json:"id"`
			Title       string `xml:"title" json:"title"`
			Price       string `xml:"price" json:"price"`
			Description string `xml:"description" json:"description"`
		} `xml:"item" json:"products"`
	} `xml:"channel" json:"channel"`
}

func LoadJSON() (map[string]product, error) {
	// map of string -> product
	products := make(map[string]product)

	xmlData, err := ioutil.ReadFile("products.xml")
	if err != nil {
		return nil, err
	}
	data := &DataFormat{}

	err = xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return nil, err
	}

	for _, prod := range data.Rss.Channel {
		products[prod.Id] = product{ID: prod.Id, Title: prod.Title, Price: prod.Price, Description: prod.Description}
	}

	return products, nil
}
