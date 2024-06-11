package kata
// https://www.codewars.com/kata/559590633066759614000063/train/go

func MinMax(arr []int) [2]int {
  min, max := arr[0], arr[0]  
  
  for _, v := range arr {
    if min > v {
      min = v
    }
    if max < v {
      max = v
    }
 
  }
  return [2]int{min,max}
  
}

