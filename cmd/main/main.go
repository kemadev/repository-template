package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!")

	var unusedVar string
	file, _ := os.Open("nonexistentfile.txt")
	fmt.Println(file)
}
