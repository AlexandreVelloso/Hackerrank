package main
import "fmt"
import "math"

func squareRoot( x uint64 ) uint64{

  // calculate the square root
  x = uint64( math.Sqrt( float64(x) ) );
  return x;
}

func main() {
      var(
        // number of test cases
        num_cases int

        // interval is [A,B]
        a, b uint64

        // a square int
        n uint64;
        x uint64;
      )

      fmt.Scan( &num_cases );

      for i := 0; i < num_cases; i++ {

        // read the input
        fmt.Scanf( "%d %d", &a, &b);

        // count the number of squase integers
        count := 0

        for x = squareRoot(a); x <= squareRoot(b); x++ {
          n = x*x;

          // if a square integer is in the interval [A,B]
          if( n >= a && n <= b ){
            count++;
          }
        }

        fmt.Println( count );
      }
}
