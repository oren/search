package main

import (
	"fmt"
	"strconv"
	// "strings"
)

// slice of maps. each map hase k/v that both are strings
// type mytype []map[string]string

type product struct {
	ID          string `json:id`
	Title       string `json:title`
	Price       string `json:price`
	Description string `json:description`
}

func main() {
	Search()
}

func Search() []product {
	// map of string -> product
	products, err := LoadJSON()
	if err != nil {
		panic(err)
	}

	keywords := createKeyWords(products)

	// slice of strings
	searchTerm := []string{"usb", "4GB", "foo"}
	results := search(searchTerm, products, keywords)
	// fmt.Println(results)
	return results
}

func createKeyWords(products map[string]product) map[string][]int {
	// for each product
	// loop on all words in title and description
	// add word to map
	// map of string -> [int, int, int]
	keywords := make(map[string][]int)

	// for _, product := range products {
	// s := []string{product.Title, product.Description}
	// s2 := strings.Join(s, " ")
	// words := strings.Fields(s2)
	// fmt.Println("joined slice", words)
	// }

	keywords["usb"] = []int{1, 2, 3, 5, 6}
	keywords["3.0"] = []int{1, 2, 3}
	keywords["8GB"] = []int{1}
	keywords["4GB"] = []int{2, 5}
	keywords["2.0"] = []int{5, 6}
	keywords["12GB"] = []int{3, 6}
	keywords["Android"] = []int{4}
	keywords["phone"] = []int{4}
	keywords["Galaxy"] = []int{4}
	keywords["x"] = []int{4}
	keywords["black"] = []int{4}
	return keywords
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
	score = RankByValue(tmpScore)
	fmt.Println("score", score)
	// return the top 5 products
	for index, value := range score {
		if index == 5 {
			break
		}
		results = append(results, products[strconv.Itoa(value.Key)])
	}

	return results
}
