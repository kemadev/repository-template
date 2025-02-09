package main

import (
	"fmt"

	"github.com/google/go-github/v60/github"
)

func main() {
	fmt.Println("Hello, World!")
	github.NewClient(nil)
}
