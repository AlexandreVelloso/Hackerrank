package main;

import "fmt";

func main(){
  var(
    // size of the array
    size int;
    // array from input
    array []int;
  )

  fmt.Scan( &size );

  array = make( []int, size );
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  // ****** END OF READ INPUT ******

  // show the array in reverse order
  for i := size-1; i >= 0; i-- {
    fmt.Printf( "%d ",array[i] );
  }
  fmt.Println("");

}
