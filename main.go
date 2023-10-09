package main

import (
	"io"
	"log"
	"net/http"
)

var apiUrl = "https://jsonplaceholder.typicode.com/posts"

func main() {
	response, err := http.Get(apiUrl)
	if err != nil {
		log.Fatalln(err)
	}

	/** Read the response body on the line below */
	body, ioErr := io.ReadAll(response.Body)
	if ioErr != nil {
		log.Fatalln(err)
	}
	/** Convert the body to type string */
	responseBody := string(body)
	log.Printf(responseBody)

	/** This can also be used */
	//fmt.Println(responseBody)
}
