package main

import (
	"fmt"
	"github.com/giovannymassuia/learning-go/05-packaging/01/math"
)

func main() {
	var m = math.NewMath(1, 2)
	fmt.Println(m.Add())
}
