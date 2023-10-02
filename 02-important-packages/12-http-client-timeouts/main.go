package main

import (
	"io"
	"net/http"
	"time"
)

func main() {

	// Create a client with a timeout of 1 second
	c := http.Client{Timeout: time.Second}

	res, err := c.Get("https://httpbin.org/delay/2")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))
}
