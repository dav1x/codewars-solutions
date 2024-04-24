package kata
import "strings" 

func Solution(number int) string {

  conversions := []struct {
    value int
    digit string
  }{
    {1000, "M"},
    {900, "CM"},
    {500, "D"},
    {400, "CD"},
    {100, "C"},
    {90, "XC"},
    {50, "L"},
    {40, "XL"},
    {10, "X"},
    {9, "IX"},
    {5, "V"},
    {4, "IV"},
    {1, "I"},
  }
    
  var romanNumeral strings.Builder
  for _, conversion := range conversions {
    for number >= conversion.value {
      romanNumeral.WriteString(conversion.digit)
      number -= conversion.value
    }
  }

  return romanNumeral.String()
        
}
