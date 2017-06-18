package main;
import"fmt";

func solve( array []int ){

  var cloud = 0;
  var numJumps = 0;

  for ( cloud < len(array) -1 ) {

    // if there is only 1 cloud left
    if( cloud == len(array)-2 ){
      numJumps++;
      break;
    }

    // first look the cloud+2
    if( array[cloud+2] == 0 ){
      cloud += 2;
    }else{
      cloud += 1;
    }

    numJumps++;
  }// end for

  fmt.Println( numJumps );
}

/**
* Method to read the input
*/
func readInput(){

  // number of clouds
  var size int;

  fmt.Scan( &size );

  array := make( []int, size );

  // read the array
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  solve( array );

}// end readInput

func main(){

  readInput();

}
