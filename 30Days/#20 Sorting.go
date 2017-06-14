package main;

import "fmt";

func main(){
  var(
    // number of swaps
    numSwaps int;

    // size of the unsorted array
    size int;
    // unsorted array
    array []int;
  )

  fmt.Scan( &size );

  array = make( []int, size );
  for i:=0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  // ****** end of read input ******

  // sort the array
  for i := 0; i < size; i++ {
    for j:= 0; j < size-1; j++ {

      if( array[j] > array[j+1] ){

        // swap
        aux := array[j];
        array[j] = array[j+1];
        array[j+1] = aux;


        numSwaps++;
      }
    }

    if( numSwaps == 0 ){
      break;
    }
  }

  fmt.Printf( "Array is sorted in %d swaps.\n",numSwaps );
  fmt.Printf( "First Element: %d\n",array[0] );
  fmt.Printf( "Last Element: %d\n",array[ len(array)-1 ] );

}
