package main

import "fmt"

func main(){
  fmt.Println(isUnique("tesg"))
  fmt.Println(isUniqueBit("test"))
}



func isUnique(s string) bool {
  a := make(map[rune]bool)

  for _, v := range s{
    if ok := a[v]; ok {
      return false
    }
    a[v] = true
  }
  return true
}

func isUniqueBit(s string) bool {
  checker := 0
  for k := range s{
    v := s[k] - "a"[0]
    if (checker & (1 << v)) > 0 {
      return false
    }
    checker |= (1 << v)
  }
  return true
}
