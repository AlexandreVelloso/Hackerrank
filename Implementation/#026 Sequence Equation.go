package main

import "fmt"

func main(){
  var(
    // size of the array
    size int

    // array from input
    array []int

    // map p( 'X' ) to equivalent 'x'
    P map[int]int
  )

  // begin of read input

  fmt.Scan( &size )

  array = make( []int, size )
  for i:=0; i < size; i++{
    fmt.Scan( &array[i] )
  }

  // end of read input

  P = make( map[int]int )
  // do the mapping
  for i:= 1; i <= size; i++ {

    // this map is
    // P(i) = array[i-1]
    P[i] = array[i-1]
  }

  /**
    Input:
    3
    2 3 1

    Explanation
    1) x = 1 -> p(3) = p(p(2)) = p(p(y)), print(2)
    2) x = 2 -> p(1) = p(p(3)) = p(p(y)), print(3)
    3) x = 3 -> p(2) = p(p(1)) = p(p(y)), print(1)
  */
  x := 1
  y := 1
  for( x <= size ){

    if( P[P[y]] == x ){
      fmt.Println( y )
      x++
      y = 0
    }

    y++
  }

}
