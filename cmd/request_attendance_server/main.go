package main

import "log"

func main() {
	err := Run()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}
