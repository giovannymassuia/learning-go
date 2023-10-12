package main

// Solving go.mod replace problem (1st approach)
// 1. go mod edit -replace github.com/giovannymassuia/learning-go/05-packaging/03/math=../math
// 2. go mod tidy

import (
	"fmt"
	"github.com/giovannymassuia/learning-go/05-packaging/03/math"
)

func main() {
	var m = math.NewMath(1, 2)
	fmt.Println(m.Add())
}
