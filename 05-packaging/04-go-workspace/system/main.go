package main

// Solving go.mod replace problem (2st approach) with workspaces
// 1. In the root folder `go work init ./math ./system`
//     - This will create a go.work file in the root folder
//     - This will `go.work` is usually ignored in .gitignore
// 2. go mod tidy
//     - If running with local and remote dependencies, it will fail for the local
//       dependencies. This is because the local dependencies are not in the
//       go.mod file. To fix this, run:
//            - `go mod tidy -e` (e for exclude)

import (
	"fmt"
	"github.com/giovannymassuia/learning-go/05-packaging/04/math"
	"github.com/google/uuid"
)

func main() {
	var m = math.NewMath(1, 2)
	fmt.Println(m.Add())

	println(uuid.New().String())
}
