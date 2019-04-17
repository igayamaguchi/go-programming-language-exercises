package main

import (
  "fmt"
  "os"
  "strings"

  "golang.org/x/net/html"
)

func main() {
  fmt.Println("start")
  doc, err := html.Parse(strings.NewReader("<html><body><a href=\"test\">test</a><a href=\"test\">test</a></body></html>"))
  if err != nil {
    fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
    os.Exit(1)
  }
  for tagName, count := range countElement(nil, doc) {
    fmt.Println(tagName, count)
  }
  fmt.Println("end")
}

func countElement(m map[string]int, n *html.Node) map[string]int {
  if m == nil {
    m = make(map[string]int)
  }

  if n.Type == html.ElementNode {
    m[n.Data]++
  }

  for c := n.FirstChild; c != nil; c = c.NextSibling {
    m = countElement(m, c)
  }

  return m
}


