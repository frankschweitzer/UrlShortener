package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// redisClient := conf.RedisClient()
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", serveHome)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	response := "URL SHORTENER"
	byteResponse := []byte(response)
	w.Write(byteResponse)
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("initiating handleShorten")

	url := r.FormValue("url")
	fmt.Printf("url received: %s", url)

	// shortenedUrl :=

	response := "<insert shortened url here>"
	byteResponse := []byte(response)

	w.Write(byteResponse)
	fmt.Println("handleShorten complete")
}
