package main

import "fmt"
func main(){
  var total int = 0

  for i := 1901; i <= 2000; i++ {
    //count months
    if((i % 4 == 0) && (i % 100 != 0)) || ((i % 400 == 0) && (i % 100 == 0)){
      total += 366
    } else {
      total += 365
    }
  }
}
