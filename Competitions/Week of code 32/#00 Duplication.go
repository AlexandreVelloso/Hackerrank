package main;

import "fmt";

/**
* Do the complement of a binary number
*/
func invert( str string ) string{
  var complement = "";

  for i := 0; i < len(str); i++ {
    if( str[i] == '1' ){
      complement += "0";
    }else{
      complement += "1";
    }
  }// end for

  return complement;
}

/**
* Concat the string with their complement
*/
func duplicate( str string ) string{

  return str + invert(str);
}

func main(){
  var(
    // number of queries
    num_q int;
    // a number x from a query
    x int;
  )

  fmt.Scan( &num_q );

  // generate the string with 1024 positions
  var str = "0";
  for( len(str) < 1000 ){
    str = duplicate( str );
  }

  for i := 0; i < num_q; i++ {
    fmt.Scan( &x );

    //fmt.Println( string(str[x]) );
  }

}
