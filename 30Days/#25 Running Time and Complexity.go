package main

import "fmt"
import "math"

func squareRoot( x int ) int{

  // calculate the square root
  x = int( math.Sqrt( float64(x) ) );
  return x;
}

func main(){

  var(
    // number of test cases
    num_cases int;
    // each number of the test case
    number int;
  )

  fmt.Scan( &num_cases );

  for i := 0; i < num_cases; i++ {

    fmt.Scan( &number );

    sqr := squareRoot(number);
    count := 1;
    
    for j := 1; j <= sqr; j++ {
      if( number%j == 0 ){
        count++;
      }
    }

    // is a prime number
    if( count == 2 && number != 1){
      fmt.Println("Prime");
    }else{
      fmt.Println("Not prime");
    }
  }
}
