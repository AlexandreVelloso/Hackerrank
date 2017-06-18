package main
import "fmt"

func main() {

    var(
        size int
        vector []int
    )

    fmt.Scan( &size )
    vector = make( []int, size )

    // read the elements of the vector
    for i:=0 ; i < size; i++ {
        fmt.Scan( &vector[i] )
    }

    // size of the biggest set
    max := 0

    for x :=0; x < size; x++ {

      // counter of number that are above the vector[x]
      c_above:= 0
      // counter of number that are under the vector[x]
      c_under := 0

      // compare all the numbers
      for i:=0; i < size; i++ {

        // i don't know explain that, but it works
        if( vector[i] == vector[x]+1 ){
          c_above++
        }else if( vector[i] == vector[x]-1 ){
          c_under++
        }else if( vector[i] == vector[x] ){
          c_above++
        }
      }// end for(i)

      if( c_above > max ){
        max = c_above
      }
      if( c_under > max ){
        max = c_under
      }

    }

    fmt.Println( max )
}
