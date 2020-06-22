// Echo3 prints its command-line arguments
package main

import (
  "fmt"
  "os"
)

func main() {
  fmt.Println(strings.Join(os.Args[1:], " "))
}
