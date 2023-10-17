package main

import (
	"fmt"
	"os"
)

func main() {
	i := 1
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file-%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("Hello World")
		i++

		if i > 200 {
			break
		}
	}
}
