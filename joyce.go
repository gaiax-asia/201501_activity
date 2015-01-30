package main

import (
	"fmt"
	"time"
)

func main() {


  yrs := 0


  for y := 1900; y < 2000; y++ {
    for m := 1; m <= 12; m++ {
      t := time.Date(y, m, 1, 0, 0, 0, 0, time.UTC)
      if t.Weekday().String() == "Sunday" {
        yrs += 1
      }

    }
  }
  fmt.Println(yrs)

}
