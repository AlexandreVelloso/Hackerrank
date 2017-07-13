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
// min int
var MIN = math.MinInt32
// max int
var MAX = math.MaxInt32

func isAbaixoDiagonalPrincipal( pos position )bool{
  return ( pos.line > pos.column );
}

func isAcimaDiagonalPrincipal( pos position ) bool{
  return ( pos.line < pos.column );
}

func isAbaixoDiagonalSecundaria( pos position )bool{
  return ( (pos.line + pos.column) >= size);
}

func isAcimaDiagonalSecundaria( pos position )bool{
  return (pos.line + pos.column < size);
}

/**
* Verifica se um obstaculo esta acima da rainha, e se ele e'
* o mais perto ate entao
*/
func isAcima( pos, obstacle, obstaculoMaisPerto position ) bool{

  // testa se o obstaculo esta na mesma coluna que a rainha
  // testa se o obstaculo e' o mais proximo dela nessa coluna
  return ( obstacle.column == pos.column &&
           obstacle.line < pos.line &&
           obstacle.line < obstaculoMaisPerto.line )
}// end function isAcima

/**
* Conta quantas casas livres a rainha tem acima dela
*/
func contAcima( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MAX , MAX };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isAcima( pos, obstacles[i], obstaculoMaisPerto) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // fmt.Println( obstaculoMaisPerto )

  // if have a obstacle
  if( obstaculoMaisPerto.line != MAX ){
    return ( pos.line - obstaculoMaisPerto.line - 1 )
  }else{
    return ( pos.line )
  }
}// end function contAcima

/**
* Verifica se um obstaculo esta na diagonal superior da rainha e se
* esse obstaculo e' o meis perto ate entao
*/
func isAcimaDir( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // para saber se um obstaculo esta na mesma diagonal direita
  // a soma da linha e coluna vao ser uma constante
  var constante = pos.line + pos.column

  // Verifica se o obstaculo esta na diagonal superior direita da rainha,
  // e se ele e' o mais perto
  // e tbm se ele esta na parte de cima, e nao na de baixo
  return ( constante == (obstaculo.line + obstaculo.column) &&
           obstaculo.line > obstaculoMaisPerto.line &&
           obstaculo.column > pos.column )
}// end function isAcimaDir

/**
* Conta quantas casas livres a rainha tem para a diagonal
* superior direita
*/
func contAcimaDir( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MIN , MIN };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isAcimaDir( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have a obstacle
  if( obstaculoMaisPerto.line != MIN ){

    return ( pos.column - obstaculoMaisPerto.line )

  }else{

    if( isAbaixoDiagonalSecundaria(pos) ){

      return ( size - 1 - pos.column )
    }else{

      return ( pos.line )
    }
  }
}// end function contAcimaDir

/**
* Verifica se um obstaculo esta na direita da rainha e se ele e' o mais
* perto ate entao
*/
func isDir( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // testa se a rainha e o obstaculo estao na mesma linha
  // testa se a coluna do obstaculo e' maior do que a coluna da rainha
  // testa se o obstaculo esta mais perto do que o anterior
  return ( pos.line == obstaculo.line &&
           pos.column < obstaculo.column &&
           obstaculo.column < obstaculoMaisPerto.column )
}// end function isDir

/**
* Conta quantas casas livres a rainha tem para a direita
*/
func contDir( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MAX , MAX };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isDir( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have obstacle
  if( obstaculoMaisPerto.line != MAX ){

    return( obstaculoMaisPerto.column - 1 - pos.column )
  }else{

    return( size - 1 - pos.column )
  }
}// end function contDir


func isAbaixoDir( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // existe uma constante i-j para todos os elementos da mesma diagonal
  var constante = pos.line - pos.column

  // testa se a coluna do obstaculo e' maior do que a da rainha
  // todos os elementos dessa diagonal terao a mesma constante i - j
  // testa se o obstaculo esta mais perto do que o anterior
  return( obstaculo.column > pos.column &&
          (obstaculo.line - obstaculo.column) == constante &&
          obstaculo.column < obstaculoMaisPerto.column )
}

func contAbaixoDir( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MAX , MAX };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isAbaixoDir( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have obstacles
  if( obstaculoMaisPerto.line != MAX ){

    return ( obstaculoMaisPerto.column - 1 - pos.column )
  }else{

    if( isAcimaDiagonalPrincipal(pos) ){
      return ( size -1 - pos.column )
    }else{
      return ( size -1 - pos.line )
    }
  }

}

/**
* Verifica se existe um obstaculo abaixo da rainha, e se ele e'
* o menor ate o momento
*/
func isAbaixo( pos, obstacle, obstaculoMaisPerto position ) bool{

  // testa se o obstaculo esta na mesma coluna que a rainha
  // testa se a linha do obstaculo e' maior do que a linha da rainha
  // testa se o obstaculo e' o mais proximo dela nessa coluna
  return ( obstacle.column == pos.column &&
           obstacle.line > pos.line &&
           obstacle.line > obstaculoMaisPerto.line )
}// end function isAcima

/**
* Conta quantas casas livres existem abaixo da rainha
*/
func contAbaixo( pos position, obstacles []position ) int{

    var obstaculoMaisPerto = position{ MIN , MIN };

    // test if have a obstacle
    for i := 0; i < len( obstacles ); i++{

      if( isAbaixo( pos, obstacles[i], obstaculoMaisPerto) ){
        obstaculoMaisPerto = obstacles[i]
      }
    }

    // if have a obstacle
    if( obstaculoMaisPerto.line != MIN ){
      return ( obstaculoMaisPerto.line - pos.line - 1 )
    }else{
      return ( size - 1 - pos.line )
    }
}// end function contAbaixo

/**
* Verifica se um obstaculo esta na diagonal superior da rainha e se
* esse obstaculo e' o meis perto ate entao
*/
func isAbaixoEsq( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // para saber se um obstaculo esta na mesma diagonal direita
  // a soma da linha e coluna vao ser uma constante
  var constante = pos.line + pos.column

  // Verifica se o obstaculo esta na diagonal superior direita da rainha,
  // e se ele e' o mais perto
  // e tbm se ele esta na parte de cima, e nao na de baixo
  return ( constante == (obstaculo.line + obstaculo.column) &&
           obstaculo.line < obstaculoMaisPerto.line &&
           obstaculo.column < pos.column )
}// end function isAcimaDir

/**
* Conta quantas casas livres a rainha tem para a diagonal
* inferior direita
*/
func contAbaixoEsq( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MAX , MAX };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isAbaixoEsq( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have a obstacle
  if( obstaculoMaisPerto.line != MAX ){

    return ( obstaculoMaisPerto.line - 1 - pos.line )
  }else{

    if( isAbaixoDiagonalSecundaria(pos) ){
      return ( size - 1 -pos.line )
    }else{
      return ( pos.column )
    }

  }
}// end function contAcimaDir

/**
* Verifica se um obstaculo esta na direita da rainha e se ele e' o mais
* perto ate entao
*/
func isEsq( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // testa se a rainha e o obstaculo estao na mesma linha
  // testa se a coluna do obstaculo e' maior do que a coluna da rainha
  // testa se o obstaculo esta mais perto do que o anterior
  return ( pos.line == obstaculo.line &&
           pos.column > obstaculo.column &&
           obstaculo.column > obstaculoMaisPerto.column )
}// end function isDir

/**
* Conta quantas casas livres a rainha tem para a direita
*/
func contEsq( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MIN , MIN };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isEsq( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have obstacle
  if( obstaculoMaisPerto.line != MIN ){

    return( pos.column - 1 - obstaculoMaisPerto.column )
  }else{

    return( pos.column )
  }
}// end function contEsq

func isAcimaEsq( pos, obstaculo, obstaculoMaisPerto position ) bool{

  // existe uma constante i-j para todos os elementos da mesma diagonal
  var constante = pos.line - pos.column

  // testa se a coluna do obstaculo e' maior do que a da rainha
  // todos os elementos dessa diagonal terao a mesma constante i - j
  // testa se o obstaculo esta mais perto do que o anterior
  return( obstaculo.column < pos.column &&
          (obstaculo.line - obstaculo.column) == constante &&
          obstaculo.column > obstaculoMaisPerto.column )
}

func contAcimaEsq( pos position, obstacles []position ) int{

  var obstaculoMaisPerto = position{ MIN , MIN };

  // test if have a obstacle
  for i := 0; i < len( obstacles ); i++{

    if( isAcimaEsq( pos, obstacles[i],obstaculoMaisPerto ) ){
      obstaculoMaisPerto = obstacles[i]
    }
  }

  // if have obstacles
  if( obstaculoMaisPerto.line != MIN ){

    return ( pos.column - 1 - obstaculoMaisPerto.column )
  }else{

    if( isAcimaDiagonalPrincipal(pos) ){
      return ( pos.line )
    }else{
      return ( pos.column )
    }
  }

}

func solve( ){
  // fmt.Println( "posQueen: ",posQueen );
  // fmt.Println( "Obstaculos: ",obstacles );

  var soma = contAcima( posQueen, obstacles )
  soma += contAcimaDir( posQueen, obstacles )
  soma += contDir( posQueen, obstacles )
  soma += contAbaixoDir( posQueen, obstacles )
  soma += contAbaixo( posQueen, obstacles )
  soma += contAbaixoEsq( posQueen, obstacles )
  soma += contEsq( posQueen, obstacles )
  soma += contAcimaEsq( posQueen, obstacles )

  fmt.Println( soma )

  // fmt.Println("acima", contAcima( posQueen, obstacles ) )
  // fmt.Println("acima dir", contAcimaDir( posQueen, obstacles ) )
  // fmt.Println("dir", contDir( posQueen, obstacles ) )
  // fmt.Println("abaixo dir", contAbaixoDir( posQueen, obstacles ) )
  // fmt.Println("abaixo", contAbaixo( posQueen, obstacles ) )
  // fmt.Println("abaixo esq", contAbaixoEsq( posQueen, obstacles ) )
  // fmt.Println("acima esq", contAcimaEsq( posQueen, obstacles ) )

  // fmt.Println( contAcima(posQueen, obstacles) )
  // fmt.Println( contAcimaDir(posQueen, obstacles) )
  // fmt.Println( contDir(posQueen, obstacles) )
  // fmt.Println( contAbaixoDir(posQueen, obstacles) ) ??????????
  // fmt.Println( contAbaixo(posQueen, obstacles) )
  // fmt.Println( contAbaixoEsq(posQueen, obstacles) )
  // fmt.Println( contEsq(posQueen, obstacles) )
  // fmt.Println( contAcimaEsq(posQueen, obstacles) ) ??????????
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
