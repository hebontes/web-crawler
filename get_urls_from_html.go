package main

import (
  "strings"
  "fmt"
  "golang.org/x/net/html"
  "net/url"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
  reader := strings.NewReader(htmlBody)
  doc, err := html.Parse(reader)
  if err != nil {
    return []string{}, fmt.Errorf("GO 13, Error %w", err)
  }
  var absoluteURLs []string
  var f func(*html.Node)
  f = func(n *html.Node){
    if n.Type == html.ElementNode && n.Data == "a"{
      for _, a := range n.Attr {
        if a.Key == "href" {
          parsedURL, err := url.Parse(a.Val)
          if err != nil {
            fmt.Errorf("Error parsing URL: %w", err)
            break
          }
          if parsedURL.Scheme == ""{
            absoluteURLs = append(absoluteURLs, rawBaseURL + parsedURL.Path)
          }else{
            absoluteURLs = append(absoluteURLs,a.Val)
          }
          break
        }
      }
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
      f(c)
    }
  }
  f(doc)


  return absoluteURLs, nil
}
