package main

import "fmt"

// func that try to reduce a string
func reduce ( s string ) (string,bool){
  // auxiliar with the string 's' reduced
  var aux string
  var reduced bool

  aux = ""

  i := 0
  for( ( i < (len(s)-1 ) ) && ( s[i] != s[i+1] ) ){
    aux += string( s[i] )
    i++
  }

  if( len(s) == 0 || i == len(s) -1 ){

    reduced = false
    aux = s
  }else{

    reduced = true

    for x := i+2; x < len(s); x++ {
      aux += string( s[x] )
    }
  }

  return aux,reduced
}

func main(){

  var(
    // string that will be reduced
    s string

    // if the string was reduced
    isReduced bool
  )

  // read the input
  fmt.Scan( &s )

  // Go do not have do while, then i will start
  // isReduced with true
  isReduced = true

  // while the string was reduced
  for( isReduced ){
    s,isReduced = reduce(s)
  }

  if( s == "" ){
      fmt.Println( "Empty String" )
  }else{
    fmt.Println( s )
  }

}
