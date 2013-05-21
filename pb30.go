package main

import (
  "fmt"
  "math"
  "strconv"
)

func main(){
  power := 5.0
  sums := []int{}
  limit := int(6 * math.Pow(9.0,power))
  for i:=2;i<limit;i++{
    // Iterate through each char in the number
    num_string := fmt.Sprintf("%d", i)
    sum_chars := 0
    for _,char := range num_string {
      // For each char, do the power of 5 and add to sum_chars
      char_int,_ := strconv.Atoi(string(char))
      sum_chars += int(math.Pow(float64(char_int),power))
    }
    // Check if value of sum_chars is the same as i (thus we have a match)
    if sum_chars == i {
      sums = append(sums, i)
    }
  }
  // Add all matches together.
  tsum := 0
  for _,i := range sums{
    tsum += i
  }
  fmt.Println(tsum)
}
