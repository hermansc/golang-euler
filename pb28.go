package main

import (
  "fmt"
)

func sumOfRing(n int) int{
  return 4*(2*n+1)*(2*n+1)-12*n
}

func main(){
  // Idea inner ring is ring number 0, next ring is ring number 1.
  // So ring number 1 has numbers: 2,3,4,5,6,7,8,9 where we add to the sum
  // 3, 5, 7 and 9. Ring number 0 equals the sum 1.
  // For a 1001x1001 matrix we have 500 such rings. 
  sum := 1
  for i:=1;i<=500;i++{
    sum += sumOfRing(i)
  }
  fmt.Println(sum)
}
