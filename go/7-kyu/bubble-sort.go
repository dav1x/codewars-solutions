package kata

func BubblesortOnce(numbers []int) []int {

    newNumbers := append([]int{}, numbers...)

    for i := 1; i < len(newNumbers); i++ {
      if newNumbers[i-1] > newNumbers[i] {
        newNumbers[i], newNumbers[i-1] = newNumbers[i-1], newNumbers[i]
      }
    }
  return newNumbers
}
