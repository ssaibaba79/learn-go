package main

import (
  "fmt"
  "os"
  "hello"
  )


func main() {
  
  if len(os.Args) > 1 {
    fmt.Println(hello.SayHello(os.Args[1:]))  
  }
}