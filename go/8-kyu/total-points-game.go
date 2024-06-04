// https://www.codewars.com/kata/5bb904724c47249b10000131

package kata
import (
  "strings"
)

func Points(games []string) int {

  score := 0
  for _, game := range games { 
  
    str := strings.Split(string(game), ":") 
    switch {
      case str[0] > str[1]:
        score += 3
      case str[0] == str[1]:
        score += 1
    }
  }
  return score
}
