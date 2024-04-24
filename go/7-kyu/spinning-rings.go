package kata

func SpinningRings(innerMax, outerMax int) int {
  innerRing, outerRing := 0, 0
  // innerRing --
  // outerRing ++
  // spin
  for i := 1; ;i++ {
    if innerRing == 0 {
      innerRing = innerMax
    } else {
      innerRing--
    }
    if outerRing == outerMax {
      outerRing = 0
    } else {
      outerRing++
    }
    if innerRing == outerRing {return i}
  }
}

