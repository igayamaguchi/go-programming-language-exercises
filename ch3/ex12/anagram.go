package ex12

import "strings"

func IsAnagram(str1 string, str2 string) bool {
  if len(str1) != len(str2) {
    return false
  }

  n := len(str1)

  for i := n - 1; i >= 0; i-- {
    r := str1[i]
    index := strings.Index(str2, string(r))

    if index == -1 {
      return false
    }

    str2 = str2[:index] + str2[index + 1:]
  }

  return true
}
