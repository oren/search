package search

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type DataFormat struct {
	Rss struct {
		Channel []struct {
			Id          int    `xml:"id" json:"id"`
			Title       string `xml:"title" json:"title"`
			Price       string `xml:"price" json:"price"`
			Description string `xml:"description" json:"description"`
			Link        string `xml:"link" json:"link"`
			Imagelink   string `xml:"image_link" json:"imageLink"`
		} `xml:"item" json:"products"`
	} `xml:"channel" json:"channel"`
}

func loadXML(file string) (map[int]Product, error) {
	// map of string -> product
	products := make(map[int]Product)

	xmlData, err := ioutil.ReadFile(file)
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
		products[prod.Id] = Product{
			ID:          prod.Id,
			Title:       prod.Title,
			Price:       prod.Price,
			Description: prod.Description,
			Link:        prod.Link,
			Imagelink:   prod.Imagelink}
	}

	return products, nil
}
