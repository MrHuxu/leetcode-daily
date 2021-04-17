package utils

import (
	"log"
	"strings"
)

func ParseItemID(content []byte) string {
	for _, line := range strings.Split(string(content), "\n") {
		if strings.Contains(line, `\>\>`) {
			arr := strings.Split(line, "/")
			if len(arr) < 2 {
				log.Fatal(line)
			} else {
				return arr[len(arr)-2]
			}
		}
	}
	return ""
}

func ParseItemIDFromURL(url string) string {
	arr := strings.Split(url, "/")
	if len(arr) < 2 {
		log.Fatal(url)
	}
	return arr[len(arr)-2]
}
