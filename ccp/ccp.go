package ccp

import (
  "fmt"
  "strconv"
)

func StrToInt(val string) int {
  value, err := strconv.Atoi(val)
  if err != nil {
    return -1
  }else
  {
    return value
  }
  
}
