package main

import (
	"net/http"
)

func main() {
	address := ":8096"
	path := "./dist"
	err := http.ListenAndServe(address, http.FileServer(http.Dir(path)))
	if err != nil {
		panic(err)
	}
}
