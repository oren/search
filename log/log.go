package logger

import (
	"fmt"
	"log"
	"net/url"

	"github.com/influxdb/influxdb/client"
)

type InfluxDBConfig struct {
	User         string
	Password     string
	DatabaseName string
}

type Logger struct {
	Database string
	User     string
	Password string
}

func NewLog(config InfluxDBConfig) (*Logger, error) {
	err := ping()
	if err != nil {
		return nil, err
	}

	l := &Logger{}
	l.Database = config.DatabaseName
	l.User = config.User
	l.Password = config.Password

	return l, nil
}

func (l *Logger) getClient() *client.Client {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}

	conf := client.Config{
		URL:      *host,
		Username: l.User,
		Password: l.Password,
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
	}

	return con
}

// TODO: try 10 times before quit
func ping() error {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		return err
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		return err
	}

	dur, ver, err := con.Ping()
	if err != nil {
		return err
	}

	log.Printf("Connected to InfluxDB %v, %s", dur, ver)
	return nil
}

func (l *Logger) Install(userID, reason string) {
	measurement := "install"
	fields := map[string]interface{}{
		"user":   userID,
		"reason": reason,
	}
	l.log(measurement, fields)
}

func (l *Logger) Uninstall(userID string) {
	measurement := "uninstall"
	fields := map[string]interface{}{
		"user": userID,
	}
	l.log(measurement, fields)
}

func (l *Logger) Search(userID string, query string) {
	measurement := "search"
	fields := map[string]interface{}{
		"user":  userID,
		"query": query,
	}
	l.log(measurement, fields)
}

func (l *Logger) Click(userID string, product int) {
	measurement := "click"
	fields := map[string]interface{}{
		"user":    userID,
		"product": product,
	}
	l.log(measurement, fields)
}

func (l *Logger) log(measurement string, fields map[string]interface{}) {
	con := l.getClient()
	bps := l.getBatchPoints(measurement, fields)
	_, err := con.Write(bps)
	if err != nil {
		log.Println(err)
	}
}

func (l *Logger) getBatchPoints(measurement string, fields map[string]interface{}) client.BatchPoints {
	var (
		sampleSize = 1
		pts        = make([]client.Point, sampleSize)
	)

	pts[0] = client.Point{
		Measurement: measurement,
		Fields:      fields,
		Precision:   "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        l.Database,
		RetentionPolicy: "default",
	}

	return bps
}
