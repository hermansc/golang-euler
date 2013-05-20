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
  // Calculate all abundant numbers
  limit := 28123
  abundants := []int{}
  for i:=1;i<=limit;i++{
    divs := divisors(i)
    if sum(divs) > i{
      abundants = append(abundants, i)
    }
  }

  // Find all possible sums of abundant numbers
  canBeWrittenAsAbundant := make(map[int]bool)
  for i:=0;i<len(abundants);i++{
    for j:=i;j<len(abundants);j++{
      s := abundants[i] + abundants[j]
      if s <= limit{
        canBeWrittenAsAbundant[s] = true
      }
    }
  }

  // Find the sum of those numbers not possible to
  // write as an abundant number
  ca_sum := 0
  for k:=1;k<=limit;k++{
    if !canBeWrittenAsAbundant[k]{
      ca_sum += k
    }
  }
  fmt.Println(ca_sum)
}
