package main

import (
	"fmt"
	"net/http"
)

func main() {
	address := ":8096"
	path := "./dist"
	fmt.Println("Listening and serving HTTP on ", address, "\r\n webroot dir ", path)
	err := http.ListenAndServe(address, http.FileServer(http.Dir(path)))
	if err != nil {
		panic(err)
	}
}
