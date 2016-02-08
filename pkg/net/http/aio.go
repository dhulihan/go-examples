package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	request_url := "http://httpbin.org/post"
	data := url.Values{}
	aio_key := os.Getenv("AIO_KEY")
	if aio_key == "" {
		data.Set("aio", aio_key)
	} else {
		log.Fatal("AIO_KEY environment variable not set.")
		os.Exit(1)
	}

	fmt.Println("Starting...")
	response, err := http.PostForm(request_url, data)

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", content)
	}

}
