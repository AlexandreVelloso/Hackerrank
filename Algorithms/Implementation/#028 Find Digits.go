package main

import "fmt"
import "strconv"

func main(){
  var(
    // number of test cases
    num_cases int
    // a number to be computed
    // i need to use this number as string to work
    number string
  )

  fmt.Scan( &num_cases )

  // for each case
  for i:=0; i < num_cases; i++ {
    fmt.Scan( &number )

    // convert the number to int
    n, _ := strconv.ParseInt( number, 10, 32 )
    // how many digits divide n
    count := 0

    for x:=0; x < len(number); x++ {

      // get the number on the position(x), using his code ascii
      frac := number[x]-48

      if( frac != 0 && n % int64(frac) == 0){
        count++
      }
    }

    fmt.Println( count )
  }
}
