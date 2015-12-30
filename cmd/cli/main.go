package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/oren/search/search"
)

var (
	ConfigFile = flag.String("config", "config.json", "Config file to load")
	Config     AppConfig
	Products   *search.Products
)

type AppConfig struct {
	Search search.SearchConfig
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
}

func main() {
	Products, err := search.New(Config.Search)
	if err != nil {
		panic(err)
	}
	query := strings.Join(os.Args, " ")
	if query != "" {
		results := Products.Search(query)
		data, err := json.Marshal(results)
		if err != nil {
			fmt.Errorf("encode response: %v", err)
		}
		os.Stdout.Write(data)
	}
}
