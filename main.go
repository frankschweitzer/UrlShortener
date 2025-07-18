package main

import (
	"fmt"

	"github.com/frankschweitzer/UrlShortener/conf"
)

func main() {
	fmt.Println("Starting project")
	redisClient := conf.RedisClient()

	fmt.Println(redisClient)
}
