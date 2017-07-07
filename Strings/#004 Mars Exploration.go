package main;
import "fmt";

/**
* Count how many letters was altered
*/
func countLetters( message string ) int{
  var count = 0;

  if( message[0] != 'S' ){
    count++;
  }
  if( message[1] != 'O' ){
    count++;
  }
  if( message[2] != 'S' ){
    count++;
  }
  return( count );
}

func main (){
  var(
    // message from input
    message string;
    // count the number of altered messages
    count int;
  )

  fmt.Scan( &message );

  for i := 0; i < len(message); i+=3 {

    var m = message[ i:i+3 ];
    count += countLetters(m);
  }// end for

  fmt.Println( count );
}
