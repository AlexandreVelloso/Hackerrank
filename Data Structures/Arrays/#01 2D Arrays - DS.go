package main;

import "fmt"

/**
* Get the sum of a hourglass
*
* @param matrix matrix from input
* @param firstLine first line of the lower glass
* @param firstColumn first column of the lower glass
*/
func sumHourglass( matrix [][]int, firstLine, firstColumn int ) int {
  var sum = 0;

  // a
  sum += matrix[ firstLine ][ firstColumn ];

  // b
  sum += matrix[ firstLine ][ firstColumn+1 ];
  // c
  sum += matrix[ firstLine ][ firstColumn+2 ];

  // d
  sum += matrix[ firstLine+1 ][ firstColumn+1 ];

  // e
  sum += matrix[ firstLine+2 ][ firstColumn ];
  // f
  sum += matrix[ firstLine+2 ][ firstColumn+1 ];
  // g
  sum += matrix[ firstLine+2 ][ firstColumn+2 ];

  return sum;
}

/**
* Method to solve the problem
*/
func solve( matrix[][]int ){

  var biggest = sumHourglass( matrix, 0,  0);

  var sum int;

  // for each hourglass
  for i := 0; i <= 3; i++ {
    for j := 0; j <= 3; j++ {

      sum = sumHourglass( matrix, i, j );
      if( sum > biggest ){
        biggest = sum;
      }

    }
  }

  fmt.Printf("%d\n",biggest);

}

/**
* Method to read the problem input
*/
func readInput(){
  var(
      matrix [][]int;
  )

  matrix = make( [][]int, 6);
  for i := range matrix {
    matrix[i] = make( []int, 6 );
  }

  for i := 0; i < 6; i++ {
    for j := 0; j < 6; j++ {
      fmt.Scan( &matrix[i][j] );
    }
  }

  // solve the problem
  solve( matrix );
}

func main(){
    readInput();
}
