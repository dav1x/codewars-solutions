// https://www.codewars.com/kata/5259b20d6021e9e14c0010d4/train/go

package kata
import "strings"

func ReverseWords(str string) string {
  
	s := strings.Split(str, " ")

	for i, v := range s {
		s[i] = Reverse(string(v))
	}
  return strings.Join(s, " ")
}

func Reverse(str string) (result string) {
 
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

