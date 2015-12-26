package main

import (
	"fmt"
	"os"
	"strings"
)

type product struct {
	ID          string `json:id`
	Title       string `json:title`
	Price       string `json:price`
	Description string `json:description`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please pass search term as a string. for example - ./search \"usb 3.0\"")
		os.Exit(1)
	}

	Search(os.Args[1])
}

func Search(term string) []product {
	// map of string -> product
	products, err := LoadJSON()
	if err != nil {
		panic(err)
	}

	// map[usb:[1 2 3 5 6] stick:[1 2 3 5 6] ...
	keywords := createKeyWords(products)

	// slice of strings
	searchTerm := strings.Fields(term)
	results := search(searchTerm, products, keywords)
	fmt.Println(results)
	return results
}

// is a string exist in a slice of strings?
func contains(s []string, e string) bool {
	for _, i := range s {
		if i == e {
			return true
		}
	}
	return false
}

// investigate the empty struct approach
// https://play.golang.org/p/aF-QpfRb6I
// https: //play.golang.org/p/vRWk64JsLb
func createKeyWords(products map[string]product) map[string][]string {
	// for each product
	// loop on all words in title and description
	// add word to map
	// map of string -> [int, int, int]

	keywords := make(map[string][]string)
	for _, product := range products {
		words := strings.Fields(product.Title + " " + product.Description)
		for _, word := range words {
			if !contains(keywords[word], product.ID) {
				keywords[word] = append(keywords[word], product.ID)
			}
		}
	}

	return keywords
}

func search(searchTerm []string, products map[string]product, keywords map[string][]string) []product {
	tmpScore := make(map[string]int)
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
	score = RankByValue(tmpScore) // [{2 2} {5 2} {1 1} {3 1} {6 1}]
	// return the top 5 products
	for index, value := range score {
		if index == 5 {
			break
		}
		results = append(results, products[value.Key])
	}

	return results
}
