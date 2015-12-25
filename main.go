package main

import (
	"fmt"
	"strconv"
)

// slice of maps. each map hase k/v that both are strings
// type mytype []map[string]string

type product struct {
	ID    string `json:id`
	Title string `json:title`
	Price string `json:price`
}

func main() {
	// map of string -> product
	products := make(map[string]product)
	products["1"] = product{ID: "1", Title: "usb 3.0 8GB", Price: "5.99"}
	products["2"] = product{ID: "2", Title: "usb 3.0 4GB", Price: "3.99"}
	products["3"] = product{ID: "3", Title: "usb 3.0 12GB", Price: "8.99"}
	products["4"] = product{ID: "4", Title: "usb 2.0 8GB", Price: "2.99"}
	products["5"] = product{ID: "5", Title: "usb 2.0 4GB", Price: "1.99"}
	products["6"] = product{ID: "6", Title: "usb 2.0 12GB", Price: "7.99"}

	// map of string -> [int, int, int]
	keywords := make(map[string][]int)
	keywords["usb"] = []int{1, 2, 3, 4, 5, 6}
	keywords["3.0"] = []int{1, 2, 3}
	keywords["8GB"] = []int{1, 4}
	keywords["4GB"] = []int{2, 5}
	keywords["2.0"] = []int{4, 5, 6}
	keywords["12GB"] = []int{3, 6}

	// slice of strings
	searchTerm := []string{"usb", "4GB", "foo"}
	results := search(searchTerm, products, keywords)
	fmt.Println(results)
}

func search(searchTerm []string, products map[string]product, keywords map[string][]int) []product {
	tmpScore := make(map[int]int)
	results := []product{}
	// for each search term
	// find its slice
	// for each number in the slice, increment a scoring map
	for _, term := range searchTerm {
		for _, productNumber := range keywords[term] {
			tmpScore[productNumber] += 1
		}
	}

	score := make(PairList, len(tmpScore))
	score = RankByWordCount(tmpScore)
	// return the top 5 products
	for index, value := range score {
		if index == 5 {
			break
		}
		results = append(results, products[strconv.Itoa(value.Key)])
	}

	return results
}
