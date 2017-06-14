package main
import "fmt"

func main() {
    var(
        // size of the array
        size int;
        // sorted array
        array []int;
        // number to be searched
        num int;

        // variables for binary search
        middle int;
        low,hight int;
    )

    fmt.Scan( &num );
    fmt.Scan( &size );

    // read the array
    array = make( []int, size );
    for i := 0; i < size; i++ {
      fmt.Scan( &array[i] );
    }

    low = 0;
    hight = size-1;

    // do a binary search
    for (low <= hight ) {

      middle = low + (hight-low)/2;

      if( array[ middle ] < num){
        low = middle+1;
      }else if( array[ middle ] > num ){
        hight = middle-1;
      }else{
        fmt.Printf( "%d\n",middle );
        break;
      }
    }

}
