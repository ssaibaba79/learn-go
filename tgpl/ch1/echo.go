package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//ex 1.2
	fmt.Println(strings.Join(os.Args[1:], " "))

	// ex 1.3
	for i, v := range os.Args[1:] {
		fmt.Println(i, v)

	}
}
