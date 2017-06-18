package main;

import "fmt";

/*
* Calculate the length of the smallest stick
*/
func small( array []int ) int{
  // length of the smallest stick begin with 0
  var length = 0;

  for i := 0; i < len(array); i++ {

    if( array[i] > 0 ){

      if( length == 0 ){
        length = array[i];
      }else if( array[i] < length ){
        length = array[i];
      }
    }// end if
  }// end for

  return length;
}

/**
* Count the number of sticks in the array
*/
func countStick( array []int )int{
  var count = 0;

  for i := 0; i < len(array); i++ {
    if( array[i] > 0 ){
      count++;
    }
  }

  return count;
}

func main(){
  var(
    // array of sticks
    array []int;
    // number of sticks, size of the array
    size int;
  )

  // read the input
  fmt.Scan( &size );
  array = make( []int, size );
  // read the array
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  // end of read the input

  fmt.Println( size );

  var smallest = small(array);
  for( smallest != 0 ){

      // subtract every position by the length
      for i := 0; i < size; i++ {
        if( array[i] > 0 ){
          array[i] -= smallest;
        }
      }

      // count the number of sticks exists in the array
      var count = countStick( array );
      if( count > 0 ){
        fmt.Println( count );
      }

      smallest = small(array);
  }

}
