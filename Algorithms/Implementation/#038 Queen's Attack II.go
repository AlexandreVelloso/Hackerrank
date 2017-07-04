package main;
import"fmt";

type position struct{
  // a position I and J on the chessboard
  line,column int
}

// position I and J of the queen on the chessboard
var posQueen position;
// obstacles positions in the chessboard
var obstacles []position;
// size of the chessboard
var size int;

func mod( a, b int ) int{
  var mod = a-b;

  if( mod < 0 ){
    mod *= -1;
  }

  return mod;
}

func min( a,b int ) int{
  if( a < b ){
    return a;
  }else{
    return b;
  }
}

func isAbaixoDiagonalPrincipal( pos position )bool{
  return ( pos.line > pos.column );
}

func isAcimaDiagonalPrincipal( pos position ) bool{
  return ( pos.line < pos.column );
}

func isAbaixoDiagonalSecundaria( pos position )bool{
  return (pos.line + pos.column >= size);
}

func isAcimaDiagonalSecundaria( pos position )bool{
  return (pos.line + pos.column < size);
}

func isDiagonalPrincipal( pos position ) bool{
  return ( pos.line == pos.column );
}

func solve( ){
  fmt.Println( "posQueen: ",posQueen );
  fmt.Println( "Obstaculos: ",obstacles );


}

/**
* Read the input data
*/
func readInput(){

  // number of obstacles
  var numObstacles int;
  // variables to read the position of the queen
  var x,y int;

  fmt.Scan( &size );
  fmt.Scan( &numObstacles );
  fmt.Scan( &x );
  fmt.Scan( &y );

  posQueen = position{ size-x , y-1 };

  obstacles = make( []position, numObstacles );

  for i := 0; i < numObstacles; i++ {

    fmt.Scanf( "%d %d",&x,&y );
    obstacles[i] = position{ size-x , y-1 };
  }

  // solve the problem
  solve( );
}// end readInput

func main(){
  // read the input data
  readInput();
}
