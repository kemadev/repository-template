package dummy

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Hello, World!!")
	// usafe pattern
	foo, err := bar()
	fmt.Println(foo, err)
}
