package main

import (
  "fmt"
  "strconv"
  "strings"
)

func getPhrase(n string, dict map[string]string) string{
  v,_ := strconv.Atoi(n)
  if v < 20 {
    return dict[strconv.Itoa(v)]
  } else if v < 100 {
    return dict[string(n[0]) + "0"] + " " + dict[string(n[1])]
  } else if v < 1000{
    r := string(n[1:])
    s := dict[string(n[0])] + " hundred"
    if r != "00" {
      return s + " and " + getPhrase(r, dict)
    } else {
      return s
    }
  } else if v < 10000 {
    r := string(n[1:])
    s := dict[string(n[0])] + " thousand"
    if r != "000" {
      return s + " " + getPhrase(r, dict)
    } else {
      return s
    }
  } else {
    return "Too big number"
  }
}

func main(){
  dictionary := make(map[string]string)
  dictionary["1"] = "one"
  dictionary["2"] = "two"
  dictionary["3"] = "three"
  dictionary["4"] = "four"
  dictionary["5"] = "five"
  dictionary["6"] = "six"
  dictionary["7"] = "seven"
  dictionary["8"] = "eight"
  dictionary["9"] = "nine"
  dictionary["10"] = "ten"
  dictionary["11"] = "eleven"
  dictionary["12"] = "twelve"
  dictionary["13"] = "thirteen"
  dictionary["14"] = "fourteen"
  dictionary["15"] = "fifteen"
  dictionary["16"] = "sixteen"
  dictionary["17"] = "seventeen"
  dictionary["18"] = "eighteen"
  dictionary["19"] = "nineteen"
  dictionary["20"] = "twenty"
  dictionary["30"] = "thirty"
  dictionary["40"] = "forty"
  dictionary["50"] = "fifty"
  dictionary["60"] = "sixty"
  dictionary["70"] = "seventy"
  dictionary["80"] = "eigthy"
  dictionary["90"] = "ninety"
  dictionary["100"] = "hundred"
  dictionary["1000"] = "thousand"
  totalLetters := 0
  for i:=1;i<=1000;i++{
    phrase := getPhrase(strconv.Itoa(i), dictionary)
    letters := len(strings.Replace(phrase, " ", "", -1))
    totalLetters += letters
  }
  fmt.Println(totalLetters)
}
