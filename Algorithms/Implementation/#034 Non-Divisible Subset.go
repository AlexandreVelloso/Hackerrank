package main;

import "fmt";

func max ( a,b int )int{
  if( a > b ){
    return a
  }else{
    return b
  }
}

func main(){

  // size of the array
  var size int
  // constant k from input
  var k int
  // array from input
  var array []int
  // array of the numbers reamainders
  var remainder []int

  // *** READ INPUT ***

  fmt.Scanf( "%d %d", &size, &k )

  array = make( []int, size )
  remainder = make( []int, size )

  for i := range( array ) {
    fmt.Scan( &array[i] )
  }

  // *** END READ INPUT ***

  count := make( []int, k )

  for i,element := range(array) {
    remainder[i] = element % k;

    count[ element%k ]++

     fmt.Println( remainder[i] )
  }

   fmt.Println("*******")
   fmt.Println( count )

  // if K is a even number
  if( k % 2 != 0 ){

    num := 0

    for i := 1; i < k/2+1; i++ {
      // fmt.Println("nums", i, k-i)
      num += max( count[ i ] , count[ k-i ] )
    }

    // if there is 1 divisible by k
    if( count[0] > 0 ){
      num += 1
    }

    fmt.Println( num )
  }else{

    num := 0

    for i := 0; i < k/2; i++ {
      fmt.Println("nums", i, k-i-1, "max:", max( count[ i ] , count[ k-i-1 ] ) )
      num += max( count[ i ] , count[ k-i-1 ] )
    }

    // if there is 1 divisible by k
    if( count[0] > 0 ){
      num += 1
    }

    fmt.Println( num, "queria 5" )
  }

  // 0 0
  // 1 2



  /*
  0 0 x
  1 0
  1 1
  1 2 x
  2 0
  2 2
  */

}
