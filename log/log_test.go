package logger

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/oren/search/log"
)

type AppConfig struct {
	InfluxDB logger.InfluxDBConfig
}

var (
	ConfigFile = "config.json"
	Config     AppConfig
	Log        *logger.Logger
)

func TestNewLog(m *testing.M) {
	ConfigBytes, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(ConfigBytes, &Config)
	if err != nil {
		log.Fatal(err)
	}

	Log, err = logger.NewLog(Config.InfluxDB)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

// products 2 and 5 must be first
// func TestSearch(t *testing.T) {
// 	products := ProductSearch.Search("usb 4GB foo")
// 	if twoAndFiveFirst(products) {
// 		return
// 	}

// 	t.Errorf("Expected ID 2 or 5 to be first, got %s", products)
// }

// func twoAndFiveFirst(products []Product) bool {
// 	if products[0].ID == 2 && products[1].ID == 5 || products[0].ID == 5 && products[1].ID == 2 {
// 		return true
// 	}

// 	return false
// }

// func BenchmarkSearch(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		ProductSearch.Search("usb 4GB foo")
// 	}
// }
