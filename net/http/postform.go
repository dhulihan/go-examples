package main

import(
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"log"
)

func main(){
	request_url := "http://httpbin.org/post"
	data := url.Values{"key": {"Value"}, "id": {"123"}}
	fmt.Println("Starting...")
	response, err := http.PostForm(request_url, data)

	if(err != nil) {
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
