package main

import (
  "fmt"
  "io"
  "net/http"
  "strings"
)

func getHTML(rawURL string)(string, error){
  res, err := http.Get(rawURL)
  if err != nil{
    fmt.Errorf("Network Response Error %w", err)
  }
  
  if res.StatusCode >= 400 {
    return "",fmt.Errorf("Error: Status Code %v", res.StatusCode) 
  }

  if !strings.Contains(res.Header.Get("content-type"), "text/html"){
    return "",fmt.Errorf("Error: conte type not text/html") 
  }
  defer res.Body.Close()

  htmlBytes,err := io.ReadAll(res.Body)
  if err !=nil{
    fmt.Errorf("Error reading io %w", err)
  }
  return string(htmlBytes),nil
}

