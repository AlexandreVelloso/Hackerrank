package main

import "fmt"

func abs( n int ) int{
  if( n < 0 ){
    return n*(-1);
  }

  return n;
}

/**
* Solve the problem
*/
func solve( matrixInput[][]int ){

  // ****** ALL MAGIC SQUARES ******
  matrix := make( [][][]int, 8 );

  matrix[0] = [][]int{ {8, 1, 6}, {3, 5, 7}, {4, 9, 2} };
  matrix[1] = [][]int{ {6, 1, 8}, {7, 5, 3}, {2, 9, 4} };
  matrix[2] = [][]int{ {4, 9, 2}, {3, 5, 7}, {8, 1, 6} };
  matrix[3] = [][]int{ {2, 9, 4}, {7, 5, 3}, {6, 1, 8} };
  matrix[4] = [][]int{ {8, 3, 4}, {1, 5, 9}, {6, 7, 2} };
  matrix[5] = [][]int{ {4, 3, 8}, {9, 5, 1}, {2, 7, 6} };
  matrix[6] = [][]int{ {6, 7, 2}, {1, 5, 9}, {8, 3, 4} };
  matrix[7] = [][]int{ {2, 7, 6}, {9, 5, 1}, {4, 3, 8} };

  // ****** END ALL MAGIC SQUARES ******

  // minimum number of changes in all squares
  var minChangeTotal = 81;

  // for each magic square
  for x := 0; x < 8; x++ {
      // minimum number of changes in this magic square
      var minChange = 0;

      // compare the matrices
      for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {

          if( matrix[x][i][j] != matrixInput[i][j] ){

            minChange += abs( matrix[x][i][j] - matrixInput[i][j] );
          }
        }
      }// end for compare

      // if this is the minimal changes
      if( minChange < minChangeTotal ){
        minChangeTotal = minChange;
      }

  }// end for each magic square

  // at the end print the minimun number of changes
  fmt.Println( minChangeTotal );

}

/**
* Read the matrix from the input data
*/
func readInput(){

  matrixInput := make( [][]int, 3 );

  for i := 0; i < 3; i++ {
    matrixInput[i] = make( []int, 3 );

    for j := 0; j < 3; j++ {
      fmt.Scan( &matrixInput[i][j] );
    }
  }

  solve( matrixInput );
}

func main(){
  readInput();
}
