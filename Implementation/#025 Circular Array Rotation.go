package main

import "fmt"

func mod ( a, b int )int{

  ret := a % b;
  if(ret < 0){
    ret+=b;
  }

  return ret;
}

func main(){
  var(
      // size of the array, n
      size int
      // number of rotations, k
      num_rotation int
      // number of queries, q
      num_query int
      // element at index to be show, m
      index int

      // index of the first element in the array
      first int

      // array from input
      array []int
  )

  fmt.Scan( &size )
  fmt.Scan( &num_rotation )
  fmt.Scan( &num_query )

  array = make( []int, size)

  // read the array
  for i:=0; i < size; i++{
    fmt.Scan( &array[i] )
  }

  // rotate the array, or just change the first
  // this is a circular array
  first = mod( 0 - num_rotation, size )

  for i:=0; i < num_query; i++{
    fmt.Scan( &index )

    // show the caractere in the position 'index'
    fmt.Println( array[ (index+first)%size ])
  }

}
