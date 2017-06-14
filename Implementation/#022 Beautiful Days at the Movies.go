package main

import "fmt"

func reverse( num int ) int{
  r := 0

  for( num > 0 ){
    r += num%10

    if( num > 9 ){
      r *= 10
    }

    num /= 10
  }

  return r
}

func abs ( num int ) int{
  if( num < 0 ){
    num *= -1
  }

  return num
}

func main (){
  var(
    // start of the interval
    interval_begin int
    // end of the interval
    interval_end int

    // const
    k int

    // number of beautiful days
    num_days int
  )

  fmt.Scan( &interval_begin )
  fmt.Scan( &interval_end )
  fmt.Scan( &k )

  num_days = 0
  for i := interval_begin; i <= interval_end; i++ {

    if( abs( i - reverse(i) )%k == 0 ){
      num_days ++
    }

  }

  fmt.Println( num_days )

}
