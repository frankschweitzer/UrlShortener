package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// redisClient := conf.RedisClient()
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", handleRedirect)

	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handling shorten")
	response := []byte{'s', 'h', 'o', 'r', 't', 'e', 'n'}
	fmt.Printf("shortened url: %s", string(response))
	w.Write(response)
}

func handleRedirect(w http.ResponseWriter, r *http.Request) {

}
