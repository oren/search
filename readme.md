# search

* [x] load xml to memory
* [x] implement search function
* [x] expose as http endpoint

## run

```
go build
./search
curl localhost:3000/search?q=usb%204GB%20foo

[{"ID":"5","Title":"usb 2.0 4GB","Price":"1.99","Description":"usb stick 2.0 4GB"},{"ID":"2","Title":"usb 3.0 4GB","Price":"3.99","Description":"usb stick 3.0 4GB red"},{"ID":"6","Title":"usb 2.0 12GB","Price":"7.99","Description":"usb stick 2.0 12GB"},{"ID":"1","Title":"usb 3.0 8GB","Price":"5.99","Description":"usb stick 3.0 8GB blue"},{"ID":"3","Title":"usb 3.0 12GB","Price":"8.99","Description":"usb stick 3.0 12GB"}]
```
(%20 is space)

## test

```
go test
```

## Nice to have

* Build a package, and make two implementations: one to run using cli and one to be a web service - https://golang.org/doc/code.html#Library
