package ex10

import (
  "bytes"
)

func Comma(str string) string {
  var buf bytes.Buffer

  n := len(str)

  if n <= 3 {
    return str
  }

  for i := 0; i <= n - 1; i++ {
    buf.WriteByte(str[i])

    if i != n - 1 && (n - i) % 3 == 1 {
      buf.WriteRune(',')
    }
  }
  return buf.String()
}
