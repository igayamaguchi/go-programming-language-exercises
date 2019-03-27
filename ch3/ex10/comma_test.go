package ex10_test

import "testing"
import "github.com/igayamaguchi/go-programming-language-exercises/ch3/ex10"

func TestNoComma(t *testing.T) {
  result := ex10.Comma("123")
  expected := "123"

  if result != expected {
    t.Errorf("result (%v) != expected (%v)", result, expected)
  }
}

func TestComma(t *testing.T) {
  result := ex10.Comma("1234567890")
  expected := "1,234,567,890"

  if result != expected {
    t.Errorf("result (%v) != expected (%v)", result, expected)
  }
}

func TestComma2(t *testing.T) {
  result := ex10.Comma("1234")
  expected := "1,234"

  if result != expected {
    t.Errorf("result (%v) != expected (%v)", result, expected)
  }
}
