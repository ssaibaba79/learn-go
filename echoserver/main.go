package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		fmt.Println(err)
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s%s \n", r.Method, r.Host, r.RequestURI)
	fmt.Fprintln(w, "Request Headers")
	for k, v := range r.Header {
		fmt.Fprintf(w, "%s:%s\n", k, v)
	}

	b, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()

	fmt.Fprintf(w, "%s", string(b))
}
