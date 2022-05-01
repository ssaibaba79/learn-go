package main

import (
	"fmt"
	"os"
)

func main() {
	readFiles()
}

func readFiles() {

	for _, fName := range os.Args[1:] {

		f, err := os.Open(fName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fmt.Println("processed file:", f.Name())

		// defer is call only when the function returns
		defer close(f)
	}
}

func close(f *os.File) {
	fmt.Println("closing file ", f.Name())
	f.Close()
}
