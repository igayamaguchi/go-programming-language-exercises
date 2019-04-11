package sample

import (
  "encoding/json"
  "fmt"
  "log"
  "os"
  "text/template"
  "time"
)

type Employee struct {
  ID        int
  Name      string
  Address   string
  DoB       time.Time
  Position  string
  Salary    int
  ManagerID int
}

type A struct {
  ID int
}

type B struct {
  ID int
}

var dilbert Employee
var employeeOfTheMonth Employee = dilbert

type Point struct{ X, Y int }

type User struct {
  ID   int    `json:"id"`
  Name string `json:",omitempty"`
}

type Article struct {
  User
  ID    string
  Title string
}

func main() {
  //employeeOfTheMonth.Position += "(proactive team player)"
  //fmt.Printf("%v\n", dilbert)
  //fmt.Printf("%v\n", employeeOfTheMonth)
  //
  //s := sample.Sample{}
  //fmt.Printf("%v", s)
  //
  //p1 := Point{1, 2}
  //p2 := Point{X: 1, Y: 2}
  //
  //fmt.Printf("%v", p1)
  //fmt.Printf("%v", p2)
  //
  //s1 := sample.Sample{}
  //s1.SetName("1")
  //s2 := sample.Sample{}
  //s2.SetName("2")
  //fmt.Printf("%v %v %v", s1.GetName(), s2.GetName(), s1 == s2)
  //
  //s11 := sample.Sample{}
  //s11.SetName("1")
  //fmt.Printf("%v", s1 == s11)
  //
  //a := Article{ID: "1", Title: "book", User: User{ID: 1, Name: "Taro"}}
  //fmt.Printf("", a.User.Name, a.Name)

  u := User{ID: 1}
  s, err := json.Marshal(u)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Println(string(s))

  u2 := User{}
  if err = json.Unmarshal(s, &u2); err != nil {
    fmt.Println(err)
  }
  fmt.Println(u2)

  const templ = `Name is {{.Name}}`

  r := User{Name: "<p>name</p>"}
  var report = template.Must(template.New("issuelist").
    Parse(templ))
  if err := report.Execute(os.Stdout, r); err != nil {
    log.Fatal(err)
  }

}
