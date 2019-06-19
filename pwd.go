package main

import (
	"log"
	"os"
)

func pwd() string {

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return dir
}
