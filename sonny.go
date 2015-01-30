package main

import (
    "fmt"
    "time"
)

func main() {
    var sundaysCount int
    t := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC)
    for years := 1; years < 101; years++ {
        for months := 1; months < 13; months++ {
            if(t.Weekday().String() == "Sunday") {
                sundaysCount++
            }
            fmt.Println(t)
            fmt.Println(t.Weekday())
            t = t.AddDate(0, 1, 0)
        }
    }
    fmt.Println(sundaysCount)
}
