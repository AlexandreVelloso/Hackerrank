package main;
import "fmt"
import "math"

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
// max int
var MIN = math.MinInt32

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

/**
* Conta quantas livres a rainha tem acima dela
*/
func contCima ( pos position, obstacles []position ) int{

  var posObstacle = position{ MIN , MIN };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( obstacles[i].column == pos.column && obstacles[i].line < posObstacle.line ){
      fmt.Println("hue")
      posObstacle = obstacles[i]
    }
  }

  // if have a obstacle
  if( posObstacle.line != MIN ){
    return ( pos.line - posObstacle.line )
  }else{
    return ( pos.line )
  }
}

func solve( ){
  fmt.Println( "posQueen: ",posQueen );
  fmt.Println( "Obstaculos: ",obstacles );

  fmt.Println( contCima(posQueen, obstacles) )
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
