package kata
import "strings"

func Parse(data string) []int{
  //TODO: write your code here
  value := 0
  returnArray := []int{}
  
  for _, letter := range strings.ToLower(data) {
    switch string(letter) {
      case "i":
        value += 1
      case "d":
        value -= 1
      case "s":
        value = value * value
      case "o":
        returnArray = append(returnArray, value)
    } 
  }
  return returnArray
}
