package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"

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
	InfluxDB logger.InfluxDBConfig
	Search   search.SearchConfig
}

// panic only doring init!
func init() {
	flag.Parse()

	ConfigBytes, err := ioutil.ReadFile(*ConfigFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(ConfigBytes, &Config)
	if err != nil {
		panic(err)
		// TODO: add line numbers to log so i can use log.fatal
		// https://golang.org/pkg/log/#pkg-examples
		// logger := log.New(os.Stderr, "OH NO AN ERROR", log.Llongfile)
	}

	Products, err = search.New(Config.Search)
	if err != nil {
		panic(err)
	}

	Log, err = logger.NewLog(Config.InfluxDB)
	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/install", func(w http.ResponseWriter, r *http.Request) {
		// TODO: if user id was not passed, generate one and return it
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Println("%s", err)
		}

		w.WriteHeader(http.StatusOK)
		s := string(out[:36])
		Log.Install(s)
		log.Println("install route")
	})

	http.HandleFunc("/uninstall", func(w http.ResponseWriter, r *http.Request) {
		// TODO: pass user id if it was passed
		w.WriteHeader(http.StatusOK)
		Log.Uninstall("323")
		log.Println("uninstall route")
	})

	http.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		// TODO: pass user id and products if they were passed
		w.WriteHeader(http.StatusOK)
		Log.Click("323", 11)
		log.Println("click route")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// TODO: pass user id, query and results if they were passed
		query := r.URL.Query().Get("q")
		if query != "" {
			results := Products.Search(query)
			w.Header().Add("Content-Type", "application/json")
			// encode it as JSON on the response
			enc := json.NewEncoder(w)
			err := enc.Encode(results)
			log.Println("query:", query, "results:", results)
			Log.Search("323", query)
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
