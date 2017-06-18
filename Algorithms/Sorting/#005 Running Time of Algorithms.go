package main
import "fmt"

func main() {
    var(
        // size of the array
        size int;
        // array to be sorted
        array []int;

        // number of shifts
        num_shift int;
    )

    fmt.Scan( &size );

    array = make( []int, size );
    for i:=0; i<size; i++{
        fmt.Scan( &array[i] );
    }

    // end of read input data

    num_shift = 0;

    for i := 1; i < size; i++{

      num := array[i];
      j := i-1;

      for( j >= 0 && array[j] > num){

        num_shift++;
        array[j+1] = array[j];
        j--;
      }

      array[j+1] = num;

    }

    fmt.Println(num_shift);
}
