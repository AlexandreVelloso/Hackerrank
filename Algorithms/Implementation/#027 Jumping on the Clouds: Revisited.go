package main

import "fmt"

func main(){
  var(
    // number of clouds
    size int

    // size of the jump
    jump int

    // array of clouds
    // 0 is a ordinary cloud
    // 1 is a thundercloud
    array []int

    // value of the actual cloud
    cloud int
    // energy
    energy int
  )

  fmt.Scan( &size )
  fmt.Scan( &jump )

  array = make( []int, size )
  for i:=0;i < size; i++{
    fmt.Scan( &array[i] )
  }

  // initialize the cloud
  cloud = 0
  // initialize the energy
  energy = 100

  // jump to the first cloud
  cloud = ( cloud + jump )%size
  if( array[ cloud ] == 0 ){
    energy -= 1
  }else{
    energy -= 3
  }

  for( cloud != 0 ){
    // jump to the first cloud
    cloud = ( cloud + jump )%size
    if( array[ cloud ] == 0 ){
      energy -= 1
    }else{
      energy -= 3
    }
  }

  fmt.Println( energy )

}
