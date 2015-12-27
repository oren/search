package search

import (
	"os"
	"testing"
)

var ProductSearch *Products

func TestMain(m *testing.M) {
	ProductSearch, _ = New("products.xml")
	os.Exit(m.Run())
}

// products 2 and 5 must be first
func TestSearch(t *testing.T) {
	products := ProductSearch.Search("usb 4GB foo")
	if twoAndFiveFirst(products) {
		return
	}

	t.Errorf("Expected ID 2 or 5 to be first, got %s", products)
}

func twoAndFiveFirst(products []Product) bool {
	if products[0].ID == 2 && products[1].ID == 5 || products[0].ID == 5 && products[1].ID == 2 {
		return true
	}

	return false
}

func BenchmarkSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProductSearch.Search("usb 4GB foo")
	}
}
