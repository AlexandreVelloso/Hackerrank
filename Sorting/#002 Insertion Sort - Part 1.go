package main
import "fmt"

func show( array []int ){

  for i := 0; i < len(array); i++ {
    fmt.Printf("%d ",array[i]);
  }

  fmt.Println();
}

func main() {
    var(
        // size of the array
        size int;
        // array to be sorted
        array []int;
        // unsorted element
        num int;
    )

    fmt.Scan( &size );

    array = make( []int, size );
    for i:=0; i<size; i++{
        fmt.Scan( &array[i] );
    }

    // end of read input data

    num = array[ size-1 ];

    i := size-2;

    // sort the array
    for ( i >=0 && array[i] > num ) {
      array[i+1] = array[i];

      show( array );
      i--;
    }

    array[i+1] = num;
    show( array );
}
