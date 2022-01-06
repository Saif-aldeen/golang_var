package ccp

import (
  "fmt"
  "strconv"
)

func StrToInt(val string) int {
  value, err := strconv.Atoi(val)
  if err != nil {
    fmt.Println("this is invalid entry")
  }else
  {
    return value
  }
  
}
