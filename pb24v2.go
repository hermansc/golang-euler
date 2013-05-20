package main

import (
  "fmt"
)

func swap(i,j int, array []int) []int{
  temp := array[i]
  array[i] = array[j]
  array[j] = temp
  return array
}

func main(){
  // Cut-the-knot algorithm, found on the internet.
  // Much faster, but not exactly intuitive.
  t := []int{0,1,2,3,4,5,6,7,8,9}
  count := 1
  nth := 1000000

  for count < nth{
    N := len(t)
    i := N-1
    for t[i-1] >= t[i]{
      i = i - 1
    }
    j := N
    for t[j-1] <= t[i-1]{
      j = j - 1
    }
    t = swap(i-1,j-1,t)
    i++
    j = N
    for i < j{
      t = swap(i-1,j-1,t)
      i++
      j--
    }
    count++
  }
  fmt.Println(t)
}
