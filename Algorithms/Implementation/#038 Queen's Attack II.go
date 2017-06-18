package main;
import"fmt";

type position struct{
  // a position I and J on the chessboard
  posI,posJ int
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

func isAbaixoDiagonalPrincipal( )bool{
  return ( posQueen.posI > posQueen.posJ );
}

func isAcimaDiagonalPrincipal( ) bool{
  return ( posQueen.posI < posQueen.posJ );
}

func isAbaixoDiagonalSecundaria( )bool{
  return (posQueen.posI + posQueen.posJ >= size);
}

func isAcimaDiagonalSecundaria( )bool{
  return (posQueen.posI + posQueen.posJ < size);
}

func isDiagonalPrincipal( ) bool{
  return ( posQueen.posI == posQueen.posJ );
}

/**
* Solved
*/
func baixo( ) int{

  var menorModulo = (size - 1 ) - posQueen.posI + 1;
  var modulo int;

  for _,element := range obstacles{

    modulo = mod( element.posI, posQueen.posI);

    if( element.posJ == posQueen.posJ && modulo < menorModulo ){
      menorModulo = modulo;
    }
  }

  return menorModulo-1 ;
}

/**
* Solved
*/
func esq( )int{

  var menorModulo = posQueen.posJ + 1;
  var modulo int;

  for _,element := range obstacles{

    modulo = mod( element.posJ, posQueen.posJ);

    if( element.posI == posQueen.posI && modulo < menorModulo ){
      menorModulo = modulo;
    }
  }

  return posQueen.posJ - (menorModulo+1);
}

/**
* Solved
*/
func cima( ) int{

  var menorModulo = (size - 1 ) - posQueen.posI + 1;
  var modulo int;

  for _,element := range obstacles{

    modulo = mod( element.posI, posQueen.posI);

    if( element.posJ == posQueen.posJ && modulo < menorModulo ){
      menorModulo = modulo;
    }
  }

  return menorModulo - 1;
}

func diagEsq1( ) int{

  var numCasas int;

  if( isAbaixoDiagonalPrincipal() ){

    // ignorando os obstaculos
    numCasas = posQueen.posJ;
  }else{

    // ignorando os obstaculos
    numCasas = posQueen.posI;
  }

  return numCasas;
}

func diagDir1( ) int{

    if( isAbaixoDiagonalSecundaria() ){
      return (size - 1) - posQueen.posJ;
    }else{
      return posQueen.posI;
    }
}

func dir( ) int{
  // ignorando obstaculos
  return (size-1) - posQueen.posJ;
}

func diagEsq2( ) int{
  if( isAcimaDiagonalPrincipal() ){
    return (size-1) - posQueen.posJ;
  }else{
    return (size-1) - posQueen.posI;
  }
}

func diagDir2( ) int{
  if( isAcimaDiagonalSecundaria() ){
    return posQueen.posJ;
  }else{
    return 7-posQueen.posI;
  }
}

func solve( ){
  fmt.Println( "posQueen: ",posQueen );
  fmt.Println( "Obstaculos: ",obstacles );

  fmt.Println( diagEsq1()+cima()+diagDir1()+dir()+diagEsq2()+baixo()+diagDir2()+esq() );

  fmt.Println( "diagEsq",diagEsq1() );
  fmt.Println( "cima",cima() );
  fmt.Println( "diagDir",diagDir1() );
  fmt.Println( "dir", dir() );
  fmt.Println( "diagEsq2", diagEsq2() );
  fmt.Println( "baixo",baixo() );
  fmt.Println( "diagDir2", diagDir2() );
  fmt.Println( "esq", esq() );
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
