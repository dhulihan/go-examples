package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h Hello
	host := "localhost:4000"
	fmt.Println("Starting server on", host)
	err := http.ListenAndServe(host, h)
	if err != nil {
		log.Fatal(err)
	}
}

