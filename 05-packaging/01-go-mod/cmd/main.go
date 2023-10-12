package main

import (
	"fmt"
	"github.com/giovannymassuia/learning-go/05-packaging/01/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
}
