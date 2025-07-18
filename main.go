package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frankschweitzer/UrlShortener/conf"
	"github.com/frankschweitzer/UrlShortener/utils"
)

type Application struct {
	redisClient conf.RedisClient
}

func main() {
	// initialize redis client
	redisClient, err := conf.NewRedisClient()
	if err != nil {
		log.Panicf("failed to initialize redis client: %s\n", err)
		return
	}

	// initialize app
	application := &Application{
		redisClient: *redisClient,
	}

	// initialize listener and functions
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/shorten", application.handleShorten)
	http.HandleFunc("/redirect/{shortUrl}", application.handleRedirect)

	fmt.Println("starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("URL SHORTENER"))
}

func (app *Application) handleShorten(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- initiating handleShorten --")

	// obtain url from form input
	url := r.FormValue("url")
	fmt.Printf("obtained url from params: %s\n", url)

	// generate shortened url
	shortenedUrl, err := utils.GetShortenCode(url)
	if err != nil {
		log.Panicf("failed to obtain url from form input: %s\n", err)
		return
	}

	// store shortened url in redis
	app.redisClient.Set(shortenedUrl, url)

	// display shortenedUrl to user
	w.Write([]byte(shortenedUrl))
	fmt.Println("-- handleShorten complete --")
}

func (app *Application) handleRedirect(w http.ResponseWriter, r *http.Request) {
	fmt.Println("-- initiating handleRedirect --")

	// obtain shortUrl from params
	key := r.PathValue("shortUrl")
	if key == "" {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}
	fmt.Printf("obtained key from PathValue: %s\n", key)

	// get the original url with the shortened url
	url := app.redisClient.Get(key)
	if url == "" {
		http.Error(w, "shortened URL Not Found", http.StatusNotFound)
		return
	}

	// redirect to original url
	fmt.Printf("redirecting to: %s", url)
	http.Redirect(w, r, url, http.StatusPermanentRedirect)

	fmt.Println("-- handleRedirect complete --")
}
