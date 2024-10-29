package kata
// https://www.codewars.com/kata/555086d53eac039a2a000083/train/go
func LoveFunc(flower1, flower2 int) bool {

  if (flower1 % 2 == 1  && flower2 % 2  == 0 ) || (flower2 % 2 == 1  && flower1 % 2  == 0) {
      return true    
  }
  return false  
}

