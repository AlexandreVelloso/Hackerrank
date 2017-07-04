package main

import "fmt"

func main (){
  var(
    // number of test cases
    num_cases int

    // number of prisioners
    num_prisioner int
    // number of sweets
    num_sweet int
    // id of the first prisioner
    id int
  )

  fmt.Scan( &num_cases )

  for i:=0; i < num_cases; i++ {

    // read the input
    fmt.Scanf("%d %d %d", &num_prisioner, &num_sweet, &id)

    // i founded that relation on the tables below
    last := (id+num_sweet-1)%num_prisioner

    if( last == 0 ){
      fmt.Println( num_prisioner )
    }else{
      fmt.Println( last )
    }

  }
}

/* tables

N : number of prisioners
M : number of sweets
S : ID of the first prisioner
R : Result

N M S | R
5 1 1 | 1
5 2 1 | 2
5 3 1 | 3
5 4 1 | 4
5 5 1 | 5
5 6 1 | 1
5 7 1 | 2
5 8 1 | 3
5 9 1 | 4

---------------------------

N M S | R
5 1 2 | 2
5 2 2 | 3
5 3 2 | 4
5 4 2 | 5
5 5 2 | 1
5 6 2 | 2
5 7 2 | 3
5 8 2 | 4
5 9 2 | 5

*/
