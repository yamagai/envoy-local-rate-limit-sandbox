package main

import (
	"io"
	"log"
	"net/http"
	"sync"
)

var requestCount int
var mu sync.Mutex

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received...")

		mu.Lock()
		requestCount++
		mu.Unlock()

		if requestCount > 3 {
			http.Error(w, "over three request!", http.StatusTooManyRequests)
			return
		}

		_, _ = io.WriteString(w, "Hello Envoy")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
