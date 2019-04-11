package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
)

type Article struct {
  Title      string `json:"title"`
  Transcript string `json:"transcript"`
}

func request(url string) (*Article, error) {

  resp, err := http.Get(url)
  if err != nil {
    return nil, err
  }

  if resp.StatusCode != http.StatusOK {
    resp.Body.Close()
    return nil, fmt.Errorf("search query failed: %s", resp.Status)
  }

  var result Article
  if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
    resp.Body.Close()
    return nil, err
  }
  resp.Body.Close()
  return &result, nil
}

func main() {
  format := "https://xkcd.com/%v/info.0.json"
  ids := os.Args[1:]

  m := make(map[string]Article)

  for _, id := range ids {
    url := fmt.Sprintf(format, id)
    fmt.Println(url)

    result, err := request(url)
    if err != nil {
      log.Fatal(err)
    }
    m[id] = *result
  }

  str, err := json.MarshalIndent(m, "", "\t")
  if err != nil {
    log.Fatal(err)
  }

  err = ioutil.WriteFile("data.json", str, os.ModePerm)
  if err != nil {
    log.Fatal(err)
  }


}
