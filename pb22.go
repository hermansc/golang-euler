package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "sort"
)

func main(){
  b, err := ioutil.ReadFile("lib/names.txt")
  if err != nil {panic(err)}
  contents := string(b)
  names := []string{}
  score := 0
  // Parse and sort names
  for _,name := range strings.Split(contents, ","){
    names = append(names, strings.Replace(name,`"`,"",-1))
  }
  sort.Strings(names)

  // Calculate scores
  for i,name := range names {
    name_score := 0
    for _,char := range []byte(name) {
      name_score += (int(char)-64)
    }
    score += (name_score * (i+1))
  }
  fmt.Println(score)
}
