package main;

import "fmt";

func main(){

  var(
    // numer of queries
    num int;
    // string from input
    str string;
    // state of the automata
    state int;
  )

  fmt.Scan( &num );

  for i:=0; i < num; i++ {

    fmt.Scan( &str );
    state = 0;

    for i := 0; i < len(str) ; i++ {

      switch str[i] {
      case 'h':
        if( state == 0 ){
          state ++;
        }
        break;

      case 'a':
        if( state == 1 || state == 7){
          state ++;
        }
        break;

      case 'c':
        if( state == 2 ){
          state ++;
        }
        break;

      case 'k':
        if( state == 3 || state == 9){
          state ++;
        }
        break;

      case 'e':
        if( state == 4 ){
          state ++;
        }
        break;

      case 'r':
        if( state == 5 || state == 6){
          state ++;
        }
        break;

      case 'n':
        if( state == 8 ){
          state ++;
        }
        break;

        }// end case
      }// end for string

    // if is in the final state
    if( state == 10 ){
      fmt.Println( "YES" );
    }else{
      fmt.Println( "NO" );
    }

  }// end for
}
