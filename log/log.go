package logger

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/influxdb/influxdb/client"
)

type Logger struct {
	Database string
	User     string
	Password string
}

func NewLog(database, user, password string) *Logger {
	ping()
	l := &Logger{}
	l.Database = database
	l.User = user
	l.Password = password
	return l
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
func ping() {
	host, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
	}
	con, err := client.NewClient(client.Config{URL: *host})
	if err != nil {
		log.Fatal(err)
	}

	dur, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Ping! %v, %s", dur, ver)
}

func (l *Logger) Install(userID string) {
	measurement := "install"
	fields := map[string]interface{}{
		"user": userID,
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
		Time:        time.Now(),
		Precision:   "s",
	}

	bps := client.BatchPoints{
		Points:          pts,
		Database:        l.Database,
		RetentionPolicy: "default",
	}

	return bps
}
