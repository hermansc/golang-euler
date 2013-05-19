package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "strconv"
)

func high_score(triangle [101][101]int) int{
  for i:=len(triangle)-3;i>=0;i--{
    for j:=len(triangle[i])-2;j>=0;j--{
      v1 := triangle[i+1][j] + triangle[i][j]
      v2 := triangle[i+1][j+1] + triangle[i][j]
      if v1 > v2 {
        triangle[i][j] = v1
      } else {
        triangle[i][j] = v2
      }
    }
  }
  return triangle[0][0]
}

func main(){
  b, err := ioutil.ReadFile("lib/triangle.txt")
  if err != nil {panic(err)}
  contents := string(b)
  triangle := [101][101]int{}
  for i,line := range strings.Split(contents, "\n"){
    for j,val := range strings.Split(strings.TrimSpace(line), " "){
      v,_ := strconv.Atoi(val)
      triangle[i][j] = v
    }
  }
  m := high_score(triangle)
  fmt.Println(m)
}
