package kata

func FindOutlier(integers []int) int {

    //https://go.dev/tour/moretypes/15
    var odds []int
    var evens []int

    for _, number := range (integers) {
        if number % 2 != 0 {
          odds = append(odds, number)
        } else {
          evens = append(evens, number)
        }

    }
    if len(odds) > len(evens) {
        return evens[0]
      } else {
        return odds[0]
      }
  
}

