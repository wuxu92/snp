package utils

import(
  _ "fmt"
)

func ErrChk(err error) {
  if err != nil {
    panic(err)
  }
}