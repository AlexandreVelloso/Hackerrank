package main;
import "fmt";

func main() {
    var (
      size int;
      array []int;
      numSwaps = 0;
    )

    // *** READ INPUT ***

    fmt.Scan( &size );

    array = make( []int, size );

    for i := 0; i < size; i++ {
        fmt.Scan( &array[i] );
    }

    // *** END READ INPUT ***

    // *** SORT ***
    for i := 0; i < size; i++ {
      for j := 0; j < size-1; j++ {

        if( array[j] > array[j+1] ){
            // swap
            aux := array[j];
            array[j] = array[j+1];
            array[j+1] = aux;

            numSwaps++;
        }
      }
    }
    // *** END SORT ***

    fmt.Printf( "Array is sorted in %d swaps.\n",numSwaps );
    fmt.Printf( "First Element: %d\n", array[0] );
    fmt.Printf( "Last Element: %d\n", array[ size-1 ] );

}
