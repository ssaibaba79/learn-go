package hello

import (
  "fmt"
  "strings"
)
func SayHello(names []string) string {
  return fmt.Sprintf("hello %s", strings.Join(names, ","))
}