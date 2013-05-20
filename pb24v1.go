package main

import (
  "fmt"
  "math/rand"
)

func removeFromArray(array []int, index int) []int{
  // Removes element at given index from array
  removed := make([]int, len(array[:index]) + len(array[index+1:]))
  copy(removed, array[:index])
  copy(removed[len(array[:index]):], array[index+1:])
  return removed
}

func removeFromIntArray(array [][]int, index int) [][]int{
  // Removes element at given index from array
  removed := make([][]int, len(array[:index]) + len(array[index+1:]))
  copy(removed, array[:index])
  copy(removed[len(array[:index]):], array[index+1:])
  return removed
}

func permutations(input []int) [][]int{
  // Returns a list of all permutations
  perms := [][]int{}
  if len(input) == 1{
    perms = append(perms, input)
  } else {
    for i:=0;i<len(input);i++{
      arr_ex := removeFromArray(input, i)
      permutations_of_subset := permutations(arr_ex)
      for j:=0;j<len(permutations_of_subset);j++{
        p := append(permutations_of_subset[j], input[i])
        perms = append(perms, p)
      }
    }
  }
  return perms
}

func swapIndexes(array [][]int,i1,i2 int) [][]int{
  // Swap indexes in a double array
  tmp := array[i1]
  array[i1] = array[i2]
  array[i2] = tmp
  return array
}


func lexigraphicSort(input [][]int) [][]int{
  // Bubble sort of a double integer array
  // Very slow
  for k:=0;k<len(input[0]);k++{
    fmt.Println("Increasing k")
    for i:=0;i<len(input);i++{
      for j:=i+1;j<len(input);j++{
        //fmt.Printf("Checking input[%d][%d] and input[%d][%d]\n", i,k,j,k)
        l := fmt.Sprintf("%v", input[i][:k])
        t := fmt.Sprintf("%v", input[j][:k])
        if l != t{ break }
        if input[i][k] > input[j][k] {
          input = swapIndexes(input, i, j)
        }
      }
    }
  }
  return input
}

func concatenate(l1,l2 [][]int) [][]int{
  concatenated := make([][]int, len(l1)+len(l2))
  copy(concatenated, l1)
  copy(concatenated[len(l1):], l2)
  return concatenated
}

func lexigraphicQuickSort(input [][]int) [][]int{
  // Quick sort of a double integer array
  // Too slow for this exercice as well.
  if len(input) <= 1{
    return input
  }
  pivot_index := rand.Intn(len(input))
  pivot_value := input[pivot_index][0]
  array := removeFromIntArray(input, pivot_index)
  less := [][]int{}
  greater := [][]int{}
  for _,v := range array{
    if v[0] <= pivot_value {
      less = append(less, v)
    } else {
      greater = append(greater, v)
    }
  }
  piv := [][]int{input[pivot_index]}
  res := concatenate(lexigraphicQuickSort(less),piv)
  res = concatenate(res, lexigraphicQuickSort(greater))
  return res
}

func main(){
  in := []int{0,1,2,3,4,5,6,7,8,9}
  l := permutations(in)
  t := lexigraphicSort(l)
  nth := 1000000
  fmt.Println(t[nth-1])
}
