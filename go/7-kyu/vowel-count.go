package kata
import "strings"

func GetCount(str string) (count int) {
  // Enter solution here
  for _, letter := range strings.ToLower(str){
    switch letter {
    case 'a', 'e', 'i', 'o', 'u': 
      count += 1
    }
  }
  return count
}
