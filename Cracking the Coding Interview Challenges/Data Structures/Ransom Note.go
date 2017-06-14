package main;

import "fmt";
import "sort";

func min( a, b int )int{
  if( a <= b ){
    return a;
  }else{
    return b;
  }
}

func main(){
  var sizeMagazine, sizeNote int;

  fmt.Scan( &sizeMagazine );
  fmt.Scan( &sizeNote );

  var magazine = make( []string, sizeMagazine );
  var note = make( []string, sizeNote );

  // *** READ INPUT ***

  // get all phases from magazine
  for i := 0; i < sizeMagazine; i++ {
    fmt.Scan( &magazine[i] );
  }

  // get all phases from note
  for i := 0; i < sizeNote; i++ {
    fmt.Scan( &note[i] );
  }

  // *** END READ INPUT ***

  // sort the arrays
  sort.Strings( magazine );
  sort.Strings( note );

  var numEqualWords = 0;
  var naoSei = min( sizeMagazine, sizeNote );
  var x = 0;
  var y = 0;

  for( x < naoSei && y < sizeMagazine ){

    if( note[x] == magazine[y] ){
      numEqualWords++;
      x++;
      y++;
    }else{
      y++;
    }
  }

  // fmt.Println( magazine );
  // fmt.Println( note );
  // fmt.Println( numEqualWords );

  if( numEqualWords == sizeNote ){
    fmt.Println("Yes");
  }else{
    fmt.Println("No");
  }

}// end main
