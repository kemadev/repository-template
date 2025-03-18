// File created by repo-as-code, however you can still modify it as you like!
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	log.Println("Hello, World!")
	err := ioutil.WriteFile("test.txt", []byte("Hello, World!"), 0o644)
	fmt.Println(err)
	return
}
