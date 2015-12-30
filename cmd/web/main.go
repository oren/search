package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/oren/search/log"
	"github.com/oren/search/search"
)

var (
	ConfigFile = flag.String("config", "config.json", "Config file to load")
	Config     AppConfig
	Products   *search.Products
	Log        *logger.Logger
)

type AppConfig struct {
	InfluxDB *SearchConfig
	Search   *SearchConfig
}

type SearchConfig struct {
	XMLFile string
}

type InfluxDBConfig struct {
	InfluxUser     string
	InfluxPassword string
}

func init() {
	flag.Parse()
	ConfigBytes, err := ioutil.ReadFile(*ConfigFile)
	if err != nil {
		log.Fatalf("Error reading config file %s\n", err)
	}
	err = json.Unmarshal(ConfigBytes, &Config)
	if err != nil {
		log.Fatalf("Error parsing config file %s\n", err)
	}

	Products, err = search.New(Config.Search)
	if err != nil {
		panic(err)
	}

	Log = logger.NewLog("search", os.Getenv("INFLUX_USER"), os.Getenv("INFLUX_PWD"))
}

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		Log.Install("323")
		log.Println("install route")
	})

	http.HandleFunc("/uninstall", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		Log.Uninstall("323")
		log.Println("uninstall route")
	})

	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		Log.Click("323", 11)
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
			Log.Search("323", "8GB")
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
