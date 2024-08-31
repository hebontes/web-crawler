package main
import (
  "fmt"
  "net/url"
)
func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int){
  currentURL, err := url.Parse(rawCurrentURL)
  if err != nil {
    fmt.Printf("Error: couldn''t parse URL")
    return
  }
  baseURL, err := url.Parse(rawBaseURL)
  if err != nil {
    fmt.Printf("Error: couldn''t parse URL")
    return
  }
  
  // skip other websites
  if currentURL.Hostname() != baseURL.Hostname() {
    return
  }

  normalizedURL, err := normalizeURL(rawCurrentURL)
  if err != nil {
    fmt.Printf("Error - normalizedURL: %v", err)
  }
    
  // Increment if visited
  if _, visited := pages[normalizedURL]; visited {
    pages[normalizedURL]++
    return
  }

  // mark as visited
  pages[normalizedURL] = 1

  fmt.Printf("Crawling %s\n", rawCurrentURL)

  htmlBody, err := getHTML(rawCurrentURL)
  if err != nil {
    fmt.Printf("Error - getHTML: %v", err)
    return
  }

  nextURLs, err := getURLsFromHTML(htmlBody, rawBaseURL)
  if err != nil {
    fmt.Printf("Error - getURLsFromHTML: %v", err)
  }

  for _, nextURL := range nextURLs {
    crawlPage(rawBaseURL, nextURL, pages)
  }
}


