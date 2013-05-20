package main

import (
  "fmt"
  "math"
)

func dayOfWeek(year, month, day int) int{
  // Gaussian algorithm found on wikipedia
  m := math.Mod(float64(month - 3), 12.0) + 12.0 + 1
  Y := year
  if m > 10 {
    Y = year - 1
  }
  y := Y % 100
  c := (Y - y)/100
  f := int(float64(2.6)*float64(m)-float64(0.2))
  w := (day + f + y + y/4 + c/4 - 2*c) % 7
  return w
}

func main(){
  fromYear := 1901
  toYear := 2000
  tot := 0
  for year:=fromYear;year<=toYear;year++{
    for month:=1;month<=12;month++{
      if dayOfWeek(year, month, 1) == 0{
        tot += 1
      }
    }
  }
  fmt.Println(tot)
}
