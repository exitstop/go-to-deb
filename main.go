package main

import (
	"log"
	"time"
)

var (
	Version     string
	GitCommit   string
	Environment string
)

func main() {
	for {
		log.Println("hello world!")
		time.Sleep(5 * time.Second)
	}
}
