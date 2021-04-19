package main

import (
	"flag"
	"log"
)

var command = flag.String("c", "", "gen or fetch")

func main() {
	flag.Parse()

	switch *command {
	case "gen":
		gen()

	case "fetch":
		fetch()

	default:
		log.Fatal("unknown command")
	}
}
