package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// write file
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	//size, err := f.WriteString("Hello World")
	size, err := f.Write([]byte("Hello World"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("wrote %d bytes\n", size)

	_ = f.Close()

	// read file
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("read %d bytes, content: %s\n", len(file), string(file))

	// read big files in chunks
	bigFile, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	// using bufio
	reader := bufio.NewReader(bigFile)
	buffer := make([]byte, 5)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			println(err.Error())
			break
		}
		fmt.Printf("read %d bytes, content: %s\n", n, string(buffer[:n]))
	}

	// delete file
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
