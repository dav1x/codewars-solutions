package kata

import "strings"

func TwoToOne(s1 string, s2 string) string {
  result := ""
  for _, char := range strings.Split("abcdefghijklmnopqrstuvwxyz", "") {
    if strings.Contains(s1+s2, char) {
      result += char
    }
  }
  return result
 }
