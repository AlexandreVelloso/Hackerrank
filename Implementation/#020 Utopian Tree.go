package main
import "fmt"

func main() {
    var(
        // number of cases
        num_cases int
        // height of the tree
        height int

        // number of cicles to test
        num_cicles int
        isSpring bool
    )

    fmt.Scan( &num_cases )

    // for each case
    for x := 0; x < num_cases; x++ {

        height = 1
        isSpring = true

        fmt.Scan( &num_cicles )

        // do the cicles
        for i := 0; i < num_cicles; i++ {

          // in the spring the tree double the height
          if( isSpring ){
              height *= 2
              isSpring = false
          }else{
            height += 1
            isSpring = true
          }

        }

        fmt.Println( height )
    }
}
