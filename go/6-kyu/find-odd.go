package kata

func FindOdd(seq []int) int {
  res := 0
  for _, x := range seq {
    res ^= x
  }
  return res
}
