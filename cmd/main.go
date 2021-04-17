package main

import (
	"flag"
	"log"
)

var command = flag.String("c", "", "gen or fetch")
var url = flag.String("u", "", "the url of the question")
var date = flag.String("d", "", "the date of the question")

func main() {
	flag.Parse()

	switch *command {
	case "gen":
		gen(*url, *date)

	case "fetch":
		fetch()

	default:
		log.Fatal("unknown command")
	}
}
