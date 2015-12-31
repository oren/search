package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"

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

func init() {
	flag.Parse()

	ConfigBytes, err := ioutil.ReadFile(*ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ConfigBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}

	Products, err = search.New(Config.Search)
	if err != nil {
		panic(err)
	}

	Log, err = logger.NewLog(Config.InfluxDB)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/health", health)
	http.HandleFunc("/install", install)
	http.HandleFunc("/uninstall", uninstall)
	http.HandleFunc("/search", searchFunc)
	http.HandleFunc("/click", click)

	log.Println("server listening")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "ok")
}

func install(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Println("%s", err)
		}
		userID = string(out[:36])
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, userID)

	Log.Install(userID)
	log.Println("install route")
}

func uninstall(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	if userID == "" {
		userID = "0"
	}

	w.WriteHeader(http.StatusOK)

	Log.Uninstall(userID)
	log.Println("uninstall route")
}

// generate user id and return it if it was not passed in querystring
func searchFunc(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	type output struct {
		Products []search.Product `json:"products"`
		UserID   string           `json:"userID"`
	}

	var ret output

	results := Products.Search(query)

	userID := r.URL.Query().Get("id")
	if userID == "" {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Println("%s", err)
		}

		user := string(out[:36])

		ret = output{
			Products: results,
			UserID:   user,
		}
	}

	if userID != "" {
		ret = output{
			Products: results,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	// encode it as JSON on the response
	enc := json.NewEncoder(w)
	err := enc.Encode(ret)

	if err != nil {
		log.Println("encode response: %v", err)
	}

	log.Println("query:", query, "results:", ret)
	Log.Search(userID, query)
}

func click(w http.ResponseWriter, r *http.Request) {
	productStr := r.URL.Query().Get("pid")
	if productStr == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
	}

	productID, err := strconv.Atoi(productStr)
	if err != nil {
		log.Println(err)
		productID = 0
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		out, err := exec.Command("uuidgen").Output()
		if err != nil {
			log.Println("%s", err)
		}
		userID = string(out[:36])
		fmt.Fprintln(w, userID)
	}

	w.WriteHeader(http.StatusOK)
	Log.Click(userID, productID)
	log.Println("click route")
}
