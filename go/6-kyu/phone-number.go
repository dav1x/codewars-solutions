package kata

import (
  "encoding/json"
  "strings"
)

func CreatePhoneNumber(numbers [10]uint) string {

//  if callme, err := json.Marshal(numbers); err != nil {

    callme, _ json.Marshall(numbers)
    s_callme := strings.ReplaceAll(strings.Trim(string(callme), "[]"), ",", "")

    return "(" + s_callme[:3] + ")" + s_callme[3:6] + "-" + s_callme[6:]

  //}
  /*
  callme := "("
 
  for i := 0;  i <len(numbers); i++ {
    if i < 2 {
       callme += fmt.Sprint(numbers[i])
    } else if i == 2 {
      callme += fmt.Sprint(numbers[i])
      callme += ") "
    } else if i < 5 {
      fmt.Println(i)
      callme += fmt.Sprint(numbers[i])
    } else if i == 5 {
      callme += fmt.Sprint(numbers[i])
      callme += "-"
    } else { 
    callme += fmt.Sprint(numbers[i])
    }
    
  }
  return callme
  */

}
