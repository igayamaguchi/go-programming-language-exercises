package ex12_test

import (
  "github.com/igayamaguchi/go-programming-language-exercises/ch3/ex12"
  "testing"
)

func TestAnagram(t *testing.T) {
  str1 := "listen"
  str2 := "silent"

  if !ex12.IsAnagram(str1, str2) {
    t.Errorf("not anagram: %v %v", str1, str2)
  }
}

func TestNotAnagram(t *testing.T) {
  str1 := "abcde"
  str2 := "12345"

  if ex12.IsAnagram(str1, str2) {
    t.Errorf("anagram: %v %v", str1, str2)
  }
}

func TestNotAnagram2(t *testing.T) {
  str1 := "abcde"
  str2 := "abcdd"

  if ex12.IsAnagram(str1, str2) {
    t.Errorf("anagram: %v %v", str1, str2)
  }
}
