package main
import "fmt"

func main() {
  var(
    // number of days
    num_days int

    // number of people that liked
    like int
    // sum of people who liked the advertisement
    sum int
  )

  fmt.Scan( &num_days )

  sum = 0
  // first day
  like = 5/2
  sum += like

  for i := 1; i < num_days; i++ {

    // number of people who liked in another day
    like = like*3/2

    sum += like
  }

  fmt.Println( sum )
}
