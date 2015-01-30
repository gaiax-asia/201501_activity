package main
import("fmt"
)

//unfinished
func main(){
  total := 0
  year := 1901
  month := 1
  day  := 1

  fmt.Println(numberOfDays(year, 10))
}

func isLeapyear(year int) int{
   if((year % 400 == 0) || (year % 4 == 0)  && !(year % 100 == 0)){
     return 1
   }else{
     return 0 
   }
}

func numberOfDays(year int, month int) int{
  switch month{
    case 1, 3, 5, 7, 8, 10, 12:
      return 31
    case 2:
      if isLeapyear(year) == 1{
        return 28
      }else{
        return 29 
      }
    default:
      return 30
  }
}
