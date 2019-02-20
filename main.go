package main

import (
	"log"
	"os"
)

func main() {
	var user   = os.Getenv("USER")
	log.Printf("Hello world from %s.", user)
}


