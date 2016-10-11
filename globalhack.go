package main

import (
  "time"
  "fmt"
)

func Hack() {
  kickOff := time.Now()
  for {
    timeDiff := int(48 - time.Now().Sub(kickOff).Hours())
    if timeDiff <= 0 {
      break
    }
    time.Sleep(1 * time.Second)
    fmt.Printf("%d hours to go #globalhack\n", timeDiff)
  }
}

func main() {
  Hack()
}
