package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	counts := make(map[string][]int)
	files := os.Args[1:]

	if len(files) == 0 {
		computeCounts(os.Stdin, counts)
	} else {
		for _, f := range files {
			//fmt.Printf("File :%s", f)
			file, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error:%v\n", err)
				continue
			}

			computeCounts(file, counts)
			file.Close()
		}
	}

	for file, cnts := range counts {
		fmt.Printf("%s %d %d %d\n", file, cnts[0], cnts[1], cnts[2])
	}
}

func computeCounts(f *os.File, counts map[string][]int) {
	input := bufio.NewScanner(f)

	cnts := make([]int, 3)
	for input.Scan() {
		cnts[0]++
		cnts[1] += len(strings.Fields(input.Text()))
		cnts[2] += len(input.Text())
	}
	counts[f.Name()] = cnts
}
