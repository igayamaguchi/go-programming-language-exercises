package ex13_test

import (
  "github.com/igayamaguchi/go-programming-language-exercises/ch3/ex13"
  "testing"
)

func TestKB(t *testing.T) {
	if ex13.KB != 1000 {
	   t.Errorf("%v", ex13.KB)
  }
}

func TestGB(t *testing.T) {
	if ex13.GB != 1000000000 {
	   t.Errorf("%v", ex13.GB)
  }
}
