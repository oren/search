package search

import "strings"

type Product struct {
	ID          string `json:id`
	Title       string `json:title`
	Price       string `json:price`
	Description string `json:description`
}

type Products struct {
	products map[string]Product
	keywords map[string]map[string]struct{}
}

func New(file string) (*Products, error) {
	products, err := loadXML(file)
	if err != nil {
		return nil, err
	}
	p := &Products{
		products: products,
	}
	p.createKeyWords()
	return p, nil
}

func (p *Products) Search(term string) []Product {
	// slice of strings
	searchTerm := strings.Fields(term)
	results := p.search(searchTerm)
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
func (p *Products) createKeyWords() {
	// for each product
	// loop on all words in title and description
	// add word to map
	// map of string -> [int, int, int]

	p.keywords = make(map[string]map[string]struct{})
	for _, product := range p.products {
		words := strings.Fields(product.Title + " " + product.Description)
		for _, word := range words {
			if p.keywords[word] == nil {
				p.keywords[word] = make(map[string]struct{})
			}
			p.keywords[word][product.ID] = struct{}{}
		}
	}
}

func (p *Products) search(searchTerm []string) []Product {
	tmpScore := make(map[string]int)
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
	// return the top 5 products
	for index, value := range score {
		if index == 5 {
			break
		}
		results = append(results, p.products[value.Key])
	}

	return results
}
