// Fetchall fetches URLS in parallel and reports their time and sizes.
package main

import (
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "os"
  "time"
)

func main() {
  start := time.Now()
  ch := make(chan string)
  for _, url := range os.Args[1:] {
    go fetch(url, ch) // start a goroutine
  }

  // write to file code :
  f, err := os.Create("./example.txt")
  if err != nil {
    fmt.Println("err")
    return
  }

  for range os.Args[1:] {
    f.WriteString(<-ch) // write to file!
    // Reasoning: you can only take things out of the channel once, so it either prints to console or file
  }
  f.Close() // something something leaks
  // Added newline to format the text file properly
  fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan <- string) {
  start := time.Now()
  resp, err := http.Get(url)

  if err != nil {
    ch <- fmt.Sprint(err) // send to channel chan ch
    return
  }

  nbytes, err := io.Copy(ioutil.Discard, resp.Body)
  resp.Body.Close() // don't leak resources
  if err != nil {
    ch <- fmt.Sprintf("While reading %s: %v", url, err)
    return
  }
  secs := time.Since(start).Seconds()
  ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

//  Exercise 1.10
// Investigate caching by running fetchall twice in succession to see whether the reported time changes much.
// Do you get the same content each time?
// Modify fetch all to print its output to a file so it can be examined.
