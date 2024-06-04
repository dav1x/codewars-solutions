package kata

func SquareSum(numbers []int) int {
    
    sum := 0
  
    for _, val := range numbers {
      sum += val * val
    }
  
   return sum
}
