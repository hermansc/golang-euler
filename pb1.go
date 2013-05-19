package main

import (
  "fmt"
  "math"
)

func findSum(lim int) int {
  sum := 0
  for i := 0; i<lim; i++ {
    if math.Mod(float64(i),3) == 0 || math.Mod(float64(i), 5) == 0 {
      sum = sum + i
    }
  }
  return sum
}

func main() {
  sum := findSum(1000)
  fmt.Println(sum)
}
