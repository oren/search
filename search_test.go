package main

import "testing"
import "reflect"

func TestSearch(t *testing.T) {
	expected := []product{}
	expected = append(expected, product{ID: "1", Title: "usb 3.0 8GB", Price: "5.99", Description: "usb stick 3.0 8GB blue"})

	actual := Search()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
