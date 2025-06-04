package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var requestCount int

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or not loaded, continue with system variables")
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/env", envHandler)

	log.Printf("The server is running on http://localhost:%s", port)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Error when starting the server: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	requestCount++
	fmt.Fprintf(w, "Go Server is working! Total requests: %d\n", requestCount)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "Current server time: %s\n", currentTime)
}

func envHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Environment Variables:")
	for _, e := range os.Environ() {
		fmt.Fprintln(w, e)
	}
}
