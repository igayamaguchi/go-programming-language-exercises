// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
  "fmt"
  "github.com/igayamaguchi/go-programming-language-exercises/ch4/github"
  "log"
  "os"
  "time"
)

type Group struct {
  LessThanOneMonth []*github.Issue
  LessThanOneYear  []*github.Issue
  OneYearOrMore    []*github.Issue
}

func outputItems(items []*github.Issue) {
  for _, item := range items {
    fmt.Printf("#%-5d %v %9.9s %.55s\n",
      item.Number, item.CreatedAt, item.User.Login, item.Title)
  }
}

//!+
func main() {
  result, err := github.SearchIssues(os.Args[1:])
  if err != nil {
    log.Fatal(err)
  }
  fmt.Printf("%d issues:\n", result.TotalCount)

  now := time.Now()
  oneMonth := now.AddDate(0, -1, 0)
  oneYear := now.AddDate(-1, 0, 0)
  fmt.Println(now, oneMonth, oneYear)

  group := Group{}
  for _, item := range result.Items {
    switch {
    case item.CreatedAt.After(oneMonth):
      group.LessThanOneMonth = append(group.LessThanOneMonth, item)
    case item.CreatedAt.After(oneYear):
      group.LessThanOneYear = append(group.LessThanOneYear, item)
    default:
      group.OneYearOrMore = append(group.OneYearOrMore, item)
    }
  }

  fmt.Println("一カ月未満")
  outputItems(group.LessThanOneMonth)
  fmt.Println("一年未満")
  outputItems(group.LessThanOneYear)
  fmt.Println("一年以上")
  outputItems(group.OneYearOrMore)
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
