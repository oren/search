package main

import "testing"

// products 2 and 5 must be first
func TestSearch(t *testing.T) {
	products := Search()
	if twoAndFiveFirst(products) {
		return
	}

	t.Errorf("Expected ID 2 or 5 to be first, got %s", products)
}

func twoAndFiveFirst(products []product) bool {
	if products[0].ID == "2" && products[1].ID == "5" || products[0].ID == "5" && products[1].ID == "2" {
		return true
	}

	return false
}
