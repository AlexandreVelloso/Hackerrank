package main;
import"fmt";

func main(){

  // number of elements
  var size int;
  fmt.Scan( &size );

  // number of occurences of each number
  // in the array
  count := make( []int, 101 );

  // array from the input
  array := make( []int, size );

  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );

    // count the number of occurences
    count[ array[i] ]++;
  }

  var numOccur = 0;

  // get the most popular element
  for i := 0; i < 101; i++ {

    if( count[i] > numOccur ){
      numOccur = count[i];
    }
  }// end for


  fmt.Println( size-numOccur );

}// end main
