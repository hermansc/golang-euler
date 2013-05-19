package main

import (
  "fmt"
)

func isPalindrome(p string) bool{
  // Check if input string is a palindrome
  s := []byte(p)
  j := len(s)-1
  for i:=0;i<len(s)/2;i++{
    if s[i] != s[j]{
      return false
    }
    j--
  }
  return true
}

func Max(a []int) int{
  // Find max value in a slice
  m := 0
  for _,v := range a {
    if v>m{
      m=v
    }
  }
  return m
}

func maxPalindrome(lim int) int{
  palindromes := []int{}
  for i:=lim;i>lim/4;i--{
    for j:=lim;j>lim/4;j--{
      product := i*j
      if isPalindrome(fmt.Sprintf("%d", product)){
        palindromes = append(palindromes, product)
      }
    }
  }
  return Max(palindromes)
}

func main(){
  max := maxPalindrome(999)
  fmt.Println(max)
}
