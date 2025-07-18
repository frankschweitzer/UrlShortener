package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frankschweitzer/UrlShortener/conf"
	"github.com/frankschweitzer/UrlShortener/utils"
)

func main() {
	// initialize redis client
	redisClient := conf.NewRedisClient()

	// initialize listener and functions
	http.HandleFunc("/shorten", handleShorten)
	http.HandleFunc("/", serveHome)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL SHORTENER"))
}

func handleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- initiating handleShorten --")

	// obtain url from form input
	url := r.FormValue("url")

	// generate shortened url
	shortenedUrl, err := utils.GetShortenCode(url)
	if err != nil {
		fmt.Printf("failed to obtain url from form input: %s", err)
		return
	}

	// store shortened url in redis

	w.Write([]byte("shortened url"))
	fmt.Println("-- handleShorten complete --")
}
