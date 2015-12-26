package main

import "testing"
import "reflect"

func TestSearch(t *testing.T) {
	expected := []product{}
	expected = append(expected, product{ID: "2", Title: "usb 3.0 4GB", Price: "3.99", Description: "usb stick 3.0 4GB red"})
	expected = append(expected, product{ID: "5", Title: "usb 2.0 4GB", Price: "1.99", Description: "usb stick 2.0 4GB"})
	expected = append(expected, product{ID: "1", Title: "usb 3.0 8GB", Price: "5.99", Description: "usb stick 3.0 8GB blue"})
	expected = append(expected, product{ID: "3", Title: "usb 3.0 12GB", Price: "8.99", Description: "usb stick 3.0 12GB"})
	expected = append(expected, product{ID: "6", Title: "usb 2.0 12GB", Price: "7.99", Description: "usb stick 2.0 12GB"})

	actual := Search()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
