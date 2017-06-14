package main
import "fmt"

func min( a int, b int) int{
  if( a <= b ){
    return a;
  }else{
    return b;
  }
}

func main() {

  // strings from the input
  var str1, str2 string;

  //
  var hashStr1 = make( []int, 26);
  var hashStr2 = make( []int, 26);

  // *** READ INPUT ***
  fmt.Scan( &str1 );
  fmt.Scan( &str2 );
  // *** END READ INPUT ***

  // fill the hash 1
  for i := 0; i < len(str1); i++ {

    caractere := str1[i];

    hashStr1[ caractere-97 ] ++;
  }
  // fill the hash 2

  for i := 0; i < len(str2); i++ {
    caractere := str2[i];

    hashStr2[ caractere-97 ] ++;
  }

  var numEquals = 0;
  // count the number of equals caracteres
  for i := 0; i < 26; i++ {
    numEquals += min( hashStr1[i], hashStr2[i] );
  }

  fmt.Println( len(str1) + len(str2) - numEquals*2 );
}
