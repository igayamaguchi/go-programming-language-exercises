package ex11_test

import (
  "github.com/igayamaguchi/go-programming-language-exercises/ch3/ex11"
  "testing"
)

func TestCommaPlus(t *testing.T) {
  result := ex11.Comma("1234567890")
  expected := "1,234,567,890"

  if result != expected {
    t.Errorf("result (%v) != expected (%v)", result, expected)
  }
}


func TestCommaMinus(t *testing.T) {
  result := ex11.Comma("-1234.1234")
  expected := "-1,234.1234"

  if result != expected {
    t.Errorf("result (%v) != expected (%v)", result, expected)
  }
}
