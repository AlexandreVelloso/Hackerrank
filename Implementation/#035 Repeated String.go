package main;
import "fmt";

func main(){

  // *** READ INPUT ***

  // string from the input
  var str string;
  // integer n from the problem
  var n int;
  // number of a's in the string
  var countA = 0;

  fmt.Scan( &str );
  fmt.Scan( &n );

  // *** END READ INPUT ***

  // count the number of a's in the string
  for i := 0; i < len(str); i++ {

    if( str[i] == 'a' ){
      countA++;
    }
  }// end for

  var numberA int;
  numberA = n/len(str) * countA;

  // ***

  var remainder int;
  remainder = n%len(str);

  // using the remainder, count the number of a's
  for i := 0; i < remainder; i++ {
    if( str[i] == 'a' ){
      numberA++;
    }
  }

  fmt.Println( numberA );
}
/**
Explanation:

sample input 0:
aba
10

i have to count the number of a's letters int the string,
it is 2.

now i count the number of times that aba repeats in a string of size n
n = 10
len(str) = 3
10/3 = 3 remainder 1

number of a's is 3*countA that is 3*2 = 6.

And finaly i have to count the a's from the remainder, it is 1, the a's
int the position 0.

6 + 1 = 7
*/
