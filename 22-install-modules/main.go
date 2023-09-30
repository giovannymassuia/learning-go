package main

// go mod tidy
// download dependencies automatically

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {

	fmt.Println("UUID:", uuid.New())

}
