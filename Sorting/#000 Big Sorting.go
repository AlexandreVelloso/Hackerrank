package main

import "fmt"

func compare(a, b string) int{
  // return 0 if 2 strings numeric values are equal
  // return 1 if a is larger
  // retrun -1 if b is larger

  if (len(a) > len(b)){
    return 1;
  }else if (len(a) < len(b)){
    return -1;
  }else{

    if( a < b ){
      return -1;
    }else if( a > b ){
      return 1;
    }else{
      return 0;
    }
  }

}

func merge(vetor []string, comeco, meio, fim int) {

  com1 := comeco;
  com2 := meio+1;
  comAux := 0;
  vetAux := make( []string, fim-comeco+1);

  for(com1<=meio && com2<=fim){

    // vetor[com1] <= vetor[com2]
    if( compare( vetor[com1], vetor[com2] ) <= 0 ){

      vetAux[comAux] = vetor[com1];
      com1++;
    }else{

      vetAux[comAux] = vetor[com2];
      com2++;
    }
    comAux++;
  }

  for(com1<=meio){  //Caso ainda haja elementos na primeira metade

    vetAux[comAux] = vetor[com1];
    comAux++;com1++;
  }

  for(com2<=fim){   //Caso ainda haja elementos na segunda metade

    vetAux[comAux] = vetor[com2];
    comAux++;com2++;
  }

  for comAux=comeco;comAux<=fim;comAux++ {    //Move os elementos de volta para o vetor original
    vetor[comAux] = vetAux[comAux-comeco];
  }

}

func mergeSort(vetor []string, comeco, fim int){

  if (comeco < fim) {

    meio := (comeco+fim)/2;

    mergeSort(vetor, comeco, meio);
    mergeSort(vetor, meio+1, fim);
    merge(vetor, comeco, meio, fim);
  }
}

func main (){

  var(
    // size of the array
    size int;
    // array to be sorted
    array []string;
  )

  // ****** READ INPUT ******

  fmt.Scan( &size );

  array = make( []string, size );
  for i := 0; i < size; i++ {
    fmt.Scan( &array[i] );
  }

  // ****** END READ INPUT ******

  mergeSort( array, 0, len(array)-1 );

  for i := 0; i < len(array); i++ {
    fmt.Println( array[i] );
  }

}
