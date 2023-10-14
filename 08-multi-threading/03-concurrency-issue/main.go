package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

// `go run -race main.go`: Used to debug or find race conditions in the code.

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		//number++
		atomic.AddUint64(&number, 1) // Atomic operations are thread safe
		//m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Number: %d\n", number)))
	})

	http.ListenAndServe(":3000", nil)
}
