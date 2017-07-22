package main;

import "fmt";

/**
* Find the smallest element in a array
*/
func small( array []int ) (int,int) {
  var smallest = array[0];
  var pos = 0;

  for i := 1; i < len(array); i++ {
    if( array[i] > 0 && array[i] < smallest ){
      smallest = array[i];
      pos = i;
    }
  }// end for

  return smallest,pos;
}

func main(){
  var(
    // number of monsters
    num_mons int;
    // seconds that Jason attacked
    seconds int;
    // Jason hit damage
    hit int;
    // array of monsters
    monster []int;
  )

  // read the input
  fmt.Scan( &num_mons, &hit, &seconds );

  monster = make( []int, num_mons );
  for i := 0; i < num_mons; i++ {
    fmt.Scan( &monster[i] );
  }
  // end of read input

  // count the number of monsters killed
  count := 0;

  // for each atack
  for i := 0; i < seconds; i++ {

    // find the monster with the smallest life point
    small,pos := small( monster );

    // if this atack kill a monster
    if( (small - hit) <= 0 ){
      count++;
    }

    // subtract the life of the monster
    monster[pos] -= hit;

  }

  // show the number of monsters killed
  fmt.Println(count);
}
