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
    )

    fmt.Scan( &size );

    array = make( []int, size );
    for i:=0; i<size; i++{
        fmt.Scan( &array[i] );
    }

    // end of read input data

    for i := 1; i < size; i++{

      num := array[i];
      j := i-1;

      for( j >= 0 && array[j] > num){

        array[j+1] = array[j];
        j--;
      }

      array[j+1] = num;

      show( array );
    }
}
