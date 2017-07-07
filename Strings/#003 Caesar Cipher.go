package main;
import "fmt";

/**
* Verify if a character is a upper case
*/
func isUpper( c uint8 ) bool{
  return (c >= 'A' && c <= 'Z');
}

/**
* Verify is a character is lower case
*/
func isLower( c uint8 ) bool{
  return (c >= 'a' && c <='z');
}
/**
* Verify if a char c is a letter
*/
func isLetter( c uint8 ) bool {
  return ( isLower(c) || isUpper(c) );
}

/**
* Encript a character, and return his string equivalent
*/
func enc( c uint8, key uint8 ) string{

  var charEnc = c;

  if( isLetter(c) ){

    if( isLower(c) ){

      charEnc = c - 97;
      // encrypt the character
      charEnc = (charEnc+key)%26;
      charEnc += 97;

    }else{
      charEnc = c - 65;
      // encrypt the character
      charEnc = (charEnc+key)%26;
      charEnc += 65;
    }

  }

  return string(charEnc);
}

func main(){
  var(
    // length of the string
    size int;
    // string to be encrypted
    line string;
    // encryption key
    key uint8;

    // string encrypted
    str_enc = "";
  )

  fmt.Scan( &size );
  fmt.Scan( &line );
  fmt.Scan( &key );

  // encode the string
  for i := 0; i < size; i++ {
    str_enc += enc( line[i], key );
  }

  fmt.Println( str_enc );

}
