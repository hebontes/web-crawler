package main

import (
  "fmt" 
  "os"
)



func main(){
  // argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]
  // Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
  if len(argsWithoutProg) < 1 {
    fmt.Println("no website provided")
    return
  }
  if len(argsWithoutProg) > 1 {
    fmt.Println("too many arguments provided")
    return
  }

  rawBaseURL := argsWithoutProg[0]
  fmt.Println("starting crawl of: ", rawBaseURL)
  pages := make(map[string]int)

  crawlPage(rawBaseURL, rawBaseURL, pages) 

  for normalizedURL, count := range pages{
    fmt.Printf("%d - %s\n", count, normalizedURL)
  }
}




