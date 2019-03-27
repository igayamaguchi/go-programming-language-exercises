package ex11

import (
  "bytes"
  "strings"
)

func Comma(str string) string {
  pre := ""

  if str[0] == '-' {
    str = str[1:]
    pre = "-"
  }

  pointIndex := strings.Index(str, ".")

  var integer, decimal string

  if pointIndex == -1 {
    integer = str
    decimal = ""
  } else {
    integer = str[:pointIndex]
    decimal = str[pointIndex+1:]
  }

  integer = comma(integer)

  result := pre + integer

  if decimal == "" {
    return result
  }

  return result + "." + decimal
}

func comma(str string) string {
  var buf bytes.Buffer
  n := len(str)

  if n <= 3 {
    return str
  }

  for i := 0; i <= n-1; i++ {
    buf.WriteByte(str[i])

    if i != n-1 && (n-i)%3 == 1 {
      buf.WriteRune(',')
    }
  }

  return buf.String()
}
