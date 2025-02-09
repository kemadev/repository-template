package main

import (
	"fmt"

	"github.com/google/go-github/v40/github"
)

func main() {
	fmt.Println("Hello, World!")
	github.NewClient(nil)
}
