package dummy

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.Println("Hello, World!!")
	// usafe pattern
	foo, err := os.Open("foo.txt")
	fmt.Println(foo)
}
