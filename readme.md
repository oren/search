# search

## Dependencies

* Go
* InfluxDB

## run

Web service

```
cd cmd/web
cp config.json.sample config.json
go build
./web
curl localhost:3000/search?q=usb%204GB%20foo

[{"ID":"5","Title":"usb 2.0 4GB","Price":"1.99","Description":"usb stick 2.0 4GB"},{"ID":"2","Title":"usb 3.0 4GB","Price":"3.99","Description":"usb stick 3.0 4GB red"},{"ID":"6","Title":"usb 2.0 12GB","Price":"7.99","Description":"usb stick 2.0 12GB"},{"ID":"1","Title":"usb 3.0 8GB","Price":"5.99","Description":"usb stick 3.0 8GB blue"},{"ID":"3","Title":"usb 3.0 12GB","Price":"8.99","Description":"usb stick 3.0 12GB"}]
```
(%20 is space)

CLI

```
cd cmd/cli
go build
./cli "usb 4GB foo"

[{"ID":"5","Title":"usb 2.0 4GB","Price":"1.99","Description":"usb stick 2.0 4GB"},{"ID":"2","Title":"usb 3.0 4GB","Price":"3.99","Description":"usb stick 3.0 4GB red"},{"ID":"6","Title":"usb 2.0 12GB","Price":"7.99","Description":"usb stick 2.0 12GB"},{"ID":"1","Title":"usb 3.0 8GB","Price":"5.99","Description":"usb stick 3.0 8GB blue"},{"ID":"3","Title":"usb 3.0 12GB","Price":"8.99","Description":"usb stick 3.0 12GB"}]
```

## test

```
cd search
go test
```

## High level

This service loads an xml of products into memory and index them. It than expose multiple HTTP endpoints: health, search, click, install, and uninstall.
The products are stored in a `map[int]Product` that looks like this: 

```
1 -> usb 3.0 sandisk 8GB 6.9
2 -> usb 3.0 sandisk 4GB 5.9
```

The keywords are stored in a `map[string]map[int]struct{}` that looks like this:

```
usb -> 1,2
3.0 -> 1,2
sandisk -> 1,2
8GB -> 1
4GB -> 2
```
