package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var requestCount int

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/env", envHandler)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprintf(w, "Go Server is working! Total requests: %d\n", requestCount)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current server time: %s", currentTime)
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	for _, e := range os.Environ() {
		fmt.Fprintln(w, e)
	}
}
