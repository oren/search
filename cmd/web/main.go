package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/oren/search"
)

var Products *search.Products
var Logger *Log

func init() {
	var err error
	Products, err = search.New("products.xml")
	if err != nil {
		panic(err)
	}

	Logger = NewLog()
}

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		Logger.Install()
		log.Println("install route")
	})

	http.HandleFunc("/uninstall", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("uninstall route")
	})

	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		log.Println("click route")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query != "" {
			results := Products.Search(query)
			w.Header().Add("Content-Type", "application/json")
			// encode it as JSON on the response
			enc := json.NewEncoder(w)
			err := enc.Encode(results)
			log.Println("query:", query, "results:", results)
			return

			// if encoding fails we log the error
			if err != nil {
				fmt.Errorf("encode response: %v", err)
			}
		}

		http.Error(w, "bad request", http.StatusBadRequest)
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
	// Search(os.Args[1])
}
