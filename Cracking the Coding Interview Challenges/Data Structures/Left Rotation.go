package main
import "fmt"

func main() {
  var size, numRotations int;
  var array[] int;

  fmt.Scanf("%d %d",&size, &numRotations);

  // *** READ INPUT ***
  array = make( []int, size );
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }
  // *** END READ INPUT ***

  // begin of the new array
  var begin = numRotations%size;

  // print the array
  for i := 0; i < size; i++ {
    fmt.Printf("%d ", array[ (begin+i)%size ] );
  }

  fmt.Println();
}
