package main

import (
  "fmt" 
  "os"
  "net/http"
  "io"
)


func getHTML(rawURL string)(string, error){

  res, err := http.Get(rawURL)
  if err != nil{
    fmt.Errorf("Network Response Error %w", err)
  }
  
  if res.StatusCode >= 400 {
    fmt.Errorf("Error: Status Code %v", res.StatusCode)
  }

  if res.Header.Get("content-type") != "text/html"{
    fmt.Errorf("Error: conte type not text/html")
  }
  defer res.Body.Close()

  htmlBytes,err := io.ReadAll(res.Body)
  if err !=nil{
    fmt.Errorf("Error reading io %w", err)
  }
  return string(htmlBytes),nil
}
func main(){
  // argsWithProg := os.Args
  argsWithoutProg := os.Args[1:]
  // Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.
  if len(argsWithoutProg) < 1 {
    fmt.Println("no website provided")
    os.Exit(1)
  }
  if len(argsWithoutProg) > 1 {
    fmt.Println("too many arguments provided")
    os.Exit(1)
  }
  BASE_URL := argsWithoutProg[0]
  fmt.Println("starting crawl of: ", BASE_URL)
  content, _ := getHTML(BASE_URL)
  fmt.Println("result: ",content)
  return
}







