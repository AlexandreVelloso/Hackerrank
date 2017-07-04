package main
import "fmt";

/**
* verify if a string is of the form
* xyxyxyxy
*/
func isValid( str string ) bool{

  if( len( str ) > 2 ){
    char1 := str[0]
    char2 := str[1]

    // char to be tested
    c := 1

    for i:= 2; i < len(str); i++ {

      // if the char to be tested is 1 -> c=1
      // and str[i] is equals to the char1
      if( c == 1 && str[i] == char1 && str[i] != char2 ){
        c = 2
      }else if( c == 2 && str[i] == char2 && str[i] != char1 ){
        c = 1
      }else{
        return false
      }
    }

    return true

  }else if( len( str ) == 2 ){

      // if the first char is equals to the second
      if( str[0] != str[1] ){
        return true
      }else{
        return false
      }
  }else{
    return false
  }
}

/**
*
*/
func twoChar( str string, char1 uint8, char2 uint8 ) string{
  var result = "";

  for i := 0; i < len(str); i++ {
    if( str[i] == char1 ){
      result += string( str[i] );
    }else if( str[i] == char2 ){
      result += string( str[i] );
    }
  }// end for

  return result;
}// emf twoChar

func main() {

  var(
    // number of elements in the string
    num int;
    // string from input
    str string;
    // existing caracteres in the string
    caracteres map[uint8]int;
  )

  fmt.Scan( &num );
  fmt.Scan( &str );

  caracteres = make( map[uint8]int );

  // mark the chars that exists in str
  for i := 0; i < num; i++ {
    caracteres[ str[i] ] = 1;
  }


  // biggest string
  max := 0

  // fazer todos os arranjos 2 a 2 com os caracteres
  // existentes na string
  for i,_ := range caracteres {
    for j,_ := range caracteres {

      // se os caracteres forem diferentes
      if( i != j ){

        // gera uma string contendo somente os 2 caracteres do arranjo
        var r = twoChar( str,i,j );

        // se essa cadeia de caracteres for valida, e maior do que a anterior
        if( isValid(r) && len(r) > max ){
          max = len(r);
        }
      }

    }
  }// fim for

  fmt.Println( max );

}
