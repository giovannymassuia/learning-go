package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}

	// defer is executed after the function ends
	defer req.Body.Close()

	fmt.Printf("Status Code: %d\n", req.StatusCode)

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Body: %s\n", res)

	// testing `defer`
	// `defer` order is LIFO (Last In First Out), basically a stack
	println("print 1")
	defer println("print 2")
	println("print 3-external-file")
	defer println("print 4-web-server")
	println("print 5-composing-templates")
}
