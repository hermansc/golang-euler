package main

import (
  "fmt"
)

func divisors(num int) []int{
  numbers := []int{}
  for i:=1;i<num;i++{
    if num%i==0{
      numbers = append(numbers, i)
    }
  }
  return numbers
}

func sum(arr []int) int{
  s := 0
  for _,val := range arr {
    s += val
  }
  return s
}

func main(){
  amicable_nums := []int{}
  for i:=1;i<10000;i++{
    sum_a := sum(divisors(i))
    sum_divs := sum(divisors(sum_a))
    if sum_divs == i{
      if i != sum_a{
        amicable_nums = append(amicable_nums, sum_a)
      }
    }
  }
  fmt.Println(sum(amicable_nums))
}
