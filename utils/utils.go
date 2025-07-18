package utils

import (
	"encoding/base64"
	"fmt"
	"time"
)

func GetShortenCode(url string) string {
	fmt.Println("initiating GetShortenCode")

	timestamp := time.Now().UnixNano()
	timestampFstring := fmt.Sprintf("%d", timestamp)
	timestampByte := []byte(timestampFstring)

	key := base64.StdEncoding.EncodeToString(timestampByte)
	key = key[:len(key)-2]

	shortenCode := key[16:]
	return shortenCode
}
