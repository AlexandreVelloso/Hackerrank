package main;

import "fmt";

/**
* Solve the staircase using recursiveness
*/
func staircaseREC( n int ) int{

  if( n == 0 || n == 1 ){
    return 1;
  }else if( n == 2 ){
    return 2;
  }else{
    return staircaseREC( n-1 ) + staircaseREC( n-2 ) + staircaseREC( n-3 );
  }

}// end method saircase

/**
* Solve the staircase using dinamic programming
*/
func staircase( array []int ) []int{

  array[0] = 1;
  array[1] = 1;
  array[2] = 2;

  for i := 3; i < len(array); i++ {
    array[i] = array[i-1] + array[i-2] + array[i-3];
  }

  return array;
}

func main(){

  var numCases int;
  var num int;

  array := make( []int, 37 );
  array = staircase( array );

  fmt.Scan( &numCases );

  for i := 0; i < numCases; i++ {
    fmt.Scan( &num );

    fmt.Println( array[num] );
  }// end for

}
