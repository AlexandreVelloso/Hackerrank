package main
import "fmt"

func main() {
    var number int

    fmt.Scan (&number )

    max := 0
    count := 0

    for( number > 0 ){

      div := number % 2

      // if the div is 0, end of the sequence
      if( div == 0 ){

        // test if the sequence is the maximum
        if( count > max ){
          max = count
        }

        count = 0
      }else{

        count ++
      }

      number /= 2
    }// end for

    // i need that test if the number finish with 1
    if( count > max ){
      max = count
    }

    fmt.Println( max )
}
