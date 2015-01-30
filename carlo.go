
//You are given the following information, but you may prefer to do some research for yourself.
//
//1 Jan 1900 was a Monday.
//Thirty days has September,
//April, June and November.
//All the rest have thirty-one,
//Saving February alone,
//Which has twenty-eight, rain or shine.
//And on leap years, twenty-nine.
//A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
//How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

package main
import ("fmt")

func is_leap_year(year int) (bool){
  // get if year is a leap year
  if year % 4  == 0 {
    if year % 100 != 0{
      return true
    }else{
      if year % 400 == 0 {
        return true
      } else {
        return false
      }
    }
  } else {
    return false
  }
}
//type list []int

func num_in_array(a int, list [4]int) (bool) {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func num_of_days(month int, year int) (int) {
  thirty_days := [4]int{9,4,6,11}
  if num_in_array(month, thirty_days) { // included in 30 array and not february
    return 30
  } else if month == 2 {
    if is_leap_year(year) {
      return 29
    }else{
      return 28
    }
  } else {
    return 31
  }
}

func main() {

  var count int = 0
  var day int = 1

  for year := 1900; year <= 2000; year++ {
    for month := 1; month <= 12; month++ {
      for date := 1; date <= num_of_days(month,year); date++ {
        if day == 7 {
          day = 1
          if date == 1 && year >= 1901{
            count++
          }
        } else{
          day++
        }
      }
    }
  }
  fmt.Println(count)
}