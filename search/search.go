package search

import (
	"path/filepath"
	"strings"
)

type SearchConfig struct {
	XMLFile string
}

type Product struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Price       string `json:"price"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Imagelink   string `json:"imageLink"`
}

type Products struct {
	products map[int]Product
	keywords map[string]map[int]struct{}
}

func New(config SearchConfig) (*Products, error) {
	products, err := loadXML(config.XMLFile)
	if err != nil {
		return nil, err
	}

	p := &Products{
		products: products,
	}
	p.createKeyWords()
	return p, nil
}

// using empty struct as a set - https://play.golang.org/p/aF-QpfRb6I
func (p *Products) createKeyWords() {
	// for each product
	//   loop on all words in title and description
	//   and k/v to each keyword

	// p.keywords is a map of strings -> map of ints
	// "usb" -> ( 5 -> {}, 1 -> {}, 3-> {} )
	// "4GB" -> (5 -> {}, 6 -> {})
	p.keywords = make(map[string]map[int]struct{})
	for _, product := range p.products {
		if validImageLink(product.Imagelink) {
			words := strings.Fields(product.Title + " " + product.Description)
			for _, word := range words {
				if p.keywords[word] == nil {
					p.keywords[word] = make(map[int]struct{})
				}
				p.keywords[word][product.ID] = struct{}{}
			}
		}
	}
}

func validImageLink(link string) bool {
	if filepath.Ext(link) == ".jpg" || filepath.Ext(link) == ".png" {
		return true
	}

	return false
}

func (p *Products) Search(term string) []Product {
	// slice of strings
	searchTerm := strings.Fields(term)
	results := p.search(searchTerm)
	return results
}

func (p *Products) search(searchTerm []string) []Product {
	tmpScore := make(map[int]int)
	results := []Product{}
	// for each search term
	// find its slice
	// for each number in the slice, increment a scoring map
	for _, term := range searchTerm {
		for productNumber := range p.keywords[term] {
			tmpScore[productNumber] += 1
		}
	}

	score := make(PairList, len(tmpScore))
	score = RankByValue(tmpScore) // [{2 2} {5 2} {1 1} {3 1} {6 1}]
	// return the top 10 products
	for index, value := range score {
		if index == 10 {
			break
		}
		results = append(results, p.products[value.Key])
	}

	return results
}
