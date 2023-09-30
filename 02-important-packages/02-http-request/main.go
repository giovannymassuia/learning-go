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
	fmt.Printf("Status Code: %d\n", req.StatusCode)

	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Body: %s\n", res)

	_ = req.Body.Close()
}
