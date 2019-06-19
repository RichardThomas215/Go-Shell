package main

import (
	"log"
	"os"
)

func rm(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}
