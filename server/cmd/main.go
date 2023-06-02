package main

import (
	"log"
)

func main() {
	_, err := db.newDatabase()
	if err != nil {
		log.Fatalf("could not init database connection: %s", err)
	}
}
