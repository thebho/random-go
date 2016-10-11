package main

import (
  "time"
  . "fmt"
)

const GH = "globalhack"

func Hack(hours int) {
  kickOff := time.Now()
  for {
    toGo := hours - int(time.Now().Sub(kickOff).Hours())
    if toGo <= 0 {
      Printf("go run %s.go #%s\n", GH, GH)
      break
    }
    Printf("%d hours to go{lang} in #%s\n", toGo, GH)
    time.Sleep(60 * 60 * time.Second)
  }
}

func main() {
  Hack(48)
}
