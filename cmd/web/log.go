package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/influxdb/influxdb/client"
)

type Log struct {
}

func NewLog() *Log {
	l := &Log{}
	return l
}

func connect() *client.Client {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	return con
}

func (l *Log) Install() {
	con := connect()

	var (
		sampleSize = 1
		pts        = make([]client.Point, sampleSize)
	)

	pts[0] = client.Point{
		Measurement: "install",
		Fields: map[string]interface{}{
			"user": "1123",
		},
		Time:      time.Now(),
		Precision: "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "search",
		RetentionPolicy: "default",
	}
	_, err := con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func Uninstall() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	var (
		sampleSize = 1
		pts        = make([]client.Point, sampleSize)
	)

	pts[0] = client.Point{
		Measurement: "uninstall",
		Fields: map[string]interface{}{
			"user": "1123",
		},
		Time:      time.Now(),
		Precision: "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "search",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func Query() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	var (
		sampleSize = 1
		pts        = make([]client.Point, sampleSize)
	)

	pts[0] = client.Point{
		Measurement: "search",
		Fields: map[string]interface{}{
			"term": "usb 3GB",
			"user": "usb 1123",
		},
		Time:      time.Now(),
		Precision: "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "search",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func Click() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: os.Getenv("INFLUX_USER"),
		Password: os.Getenv("INFLUX_PWD"),
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	var (
		sampleSize = 1
		pts        = make([]client.Point, sampleSize)
	)

	pts[0] = client.Point{
		Measurement: "click",
		Fields: map[string]interface{}{
			"product": 3,
			"user":    "usb 1123",
		},
		Time:      time.Now(),
		Precision: "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        "search",
		RetentionPolicy: "default",
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}