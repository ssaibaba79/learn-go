package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("usage curl <url>")
		os.Exit(1)
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error:%s", err)
	}

	fmt.Printf("%s \n", resp.Status)
	fmt.Printf("%s %s \n", resp.Request.Method, resp.Request.URL)
	fmt.Printf("%s \n", b)
}
