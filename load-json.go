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
			Id       string `xml:"id" json:"id"`
			Quantity int    `xml:"quantity" json:"quantity"`
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

	fmt.Println("channel", data)

	for _, prod := range data.Rss.Channel {
		products[prod.Id] = product{ID: prod.Id, Title: "usb 3.0 8GB", Price: "5.99"}
	}

	return products, nil
}
