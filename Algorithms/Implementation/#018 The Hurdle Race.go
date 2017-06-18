package main

import "fmt"

func main(){
  var (
    n,k int

    h []int

    // biggest hurdle on the track1
    max int
  )

  // number of hurdles
  fmt.Scan( &n )
  // maximum height he can jump without using potion
  fmt.Scan( &k )

  h = make( []int, n )
  for x := range h {
    fmt.Scan( &h[x] )
  }

  // the heights start in 1, then max = 0 is correct
  max = 0

  // find the biggest hurdle
  for x := range h{
    if( h[x] > max ){
      max = h[x]
    }
  }

  // print the number of potions
  if( (max - k) > 0 ){
    fmt.Println( (max - k ) )
  }else{
    fmt.Println( 0 )
  }

}
