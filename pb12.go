package main

import (
  "fmt"
  "math"
)

func removeDuplicates(a []int) []int {
  result := []int{}
  seen := map[int]int{}
  for _, val := range a {
    if _, ok := seen[val]; !ok {
      result = append(result, val)
      seen[val] = val
    }
  }
  return result
}

func findDivisors(num int) []int{
  divisors := []int{}
  // Speed optimization: Find all factors until root
  for i:=1;i<=int(math.Sqrt(float64(num)))+1;i++{
    if num%i==0{
      divisors = append(divisors, i)
    }
  }
  // Then calculate complimentary factors and add them.
  hdiv := []int{}
  for _,v := range divisors{
    if (num/v) != v{
      hdiv = append(hdiv, (num/v))
    }
  }
  for _,v := range hdiv{
    divisors = append(divisors, v)
  }
  // And clean up so we only have unique factors.
  return removeDuplicates(divisors)
}

func main(){
  lim = 500
  triangle_num := 0
  for i:=1;;i++{
    triangle_num += i
    divisors := findDivisors(triangle_num)
    if len(divisors) >= lim{
      fmt.Println(triangle_num)
      break
    }
  }
}
