// Fetch prints the content found at a URL>
package main

import (
  "fmt"
  "io"
  "net/http"
  "os"
  "strings"
)

func main() {
  pref := "http://"
  for _, url := range os.Args[1:] {
    if !strings.HasPrefix(url, pref) {
      url = pref + url
    }
    resp, err := http.Get(url)
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
      os.Exit(1)
    }
    // Implement fetch to also print out the HTTP status code (resp.Status)
    fmt.Println("Status Code is: ", resp.Status)

    b, err := io.Copy(os.Stdout, resp.Body)
    resp.Body.Close()
    if err != nil {
      fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
      os.Exit(1)
    }
    fmt.Printf("%s", b)
  }
}
