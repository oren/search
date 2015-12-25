package main

import (
	"fmt"
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

	// map of string -> [int, int, int]
	keywords := make(map[string][]int)
	keywords["usb"] = []int{1, 2}
	keywords["3.0"] = []int{1, 2}
	keywords["8GB"] = []int{1}
	keywords["4GB"] = []int{2}

	// slice of strings
	searchTerm := []string{"usb", "4GB", "foo"}
	results := search(searchTerm, products, keywords)
	fmt.Println(results)
}

func search(searchTerm []string, products map[string]product, keywords map[string][]int) []product {
	score := make(map[int]int)
	// for each search term
	// find its slice
	// for each number in the slice, increment a scoring map
	for _, term := range searchTerm {
		fmt.Println(keywords[term])
		for _, productNumber := range keywords[term] {
			fmt.Println(productNumber)
			score[productNumber] += 1
		}
	}

	score2 := make(PairList, len(score))
	score2 = RankByWordCount(score)
	fmt.Println(score2)
	// return the top 5 products
	results := []product{products["1"], products["2"]}

	return results
}
