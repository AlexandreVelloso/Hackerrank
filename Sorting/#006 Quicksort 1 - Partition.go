package main

import "fmt"

func main(){
  var(
    // size of the array
    size int;
    // unsorted array
    array []int;
    // pivot of the quicksort
    pivot int;

    left,equal,right []int;
  )

  fmt.Scan( &size );

  left = make( []int, 0);
  equal = make( []int, 0);
  right = make( []int, 0 );

  array = make( []int, size );
  // read the array from the input
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  pivot = array[0];
  for i := 0; i < size; i++ {
    if( array[i] < pivot ){
      left = append(left, array[i] );
    }else if( array[i] > pivot ){
      right = append(right, array[i] );
    }else{
      equal = append(equal, array[i] );
    }

  }

  for i := 0; i < len(left); i++ {
      fmt.Printf( "%d ",left[i] );
  }
  for i := 0; i < len(equal); i++ {
      fmt.Printf( "%d ",equal[i] );
  }
  for i := 0; i < len(right); i++ {
      fmt.Printf( "%d ",right[i] );
  }
  fmt.Println();

}
