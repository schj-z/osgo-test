package main

import (
	"log"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

func fileExists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        value = fallback
    }
    return value
}

func main() {
	if fileExists(".env") {
		loadEnv()
	} else {
		fmt.Println("No .env file found. Using system environment variables.")
	}

  bindHost := getEnv("BIND_HOST", "127.0.0.1")
  port, err := strconv.Atoi(getEnv("BIND_PORT", "8080"))
	if err != nil {
		log.Fatal("Error parsing BIND_PORT")
	}

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  	fmt.Fprintf(w, "Hello, World!")
  })

	fmt.Printf("Listening on %s:%d", bindHost, port)
  log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", bindHost, port), nil))
}