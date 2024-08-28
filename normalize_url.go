package main

import (
  "net/url"
  "fmt"
  "strings"
)

func normalizeURL(rawURL string) (string, error){
  parsedURL, err := url.Parse(rawURL)
  if err != nil {
    return "", fmt.Errorf("couldn't parse url")
  }

  fullPath:= parsedURL.Host + parsedURL.Path

  fullPath = strings.ToLower(fullPath)
  fullPath = strings.TrimSuffix(fullPath, "/")
  return fullPath, nil
}
