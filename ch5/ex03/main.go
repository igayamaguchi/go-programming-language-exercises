package main

import (
	"fmt"
	"os"
  "strings"

  "golang.org/x/net/html"
)

func main() {
  fmt.Println("start")
	doc, err := html.Parse(strings.NewReader(`<html>
<head>
<style>
</style>
</head>
<body>
<div>
  div
  <span>span</span
</div>
<script>script</script>
</body>
</html>
`))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range getText(nil, doc) {
		fmt.Println(link)
	}
  fmt.Println("end")
}

// 改行未対応
func getText(textList []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
	  textList = append(textList, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		textList = getText(textList, c)
	}
	return textList
}
