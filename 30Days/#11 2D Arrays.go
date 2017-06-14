package main
import "fmt"

func sumHourGlass ( vector [][]int, i int, j int ) (int){
    sum := 0

    /*
       sum in this order
        1 2 3
          4
        5 6 7
    */

    // 1,2 and 3
    sum += vector[i][j] + vector[i][j+1] + vector[i][j+2]
    // 4
    sum += vector[i+1][j+1]
    // 5, 6 and 7
    sum += vector[i+2][j] + vector[i+2][j+1] + vector[i+2][j+2]

    return sum
}

func main() {
    var(
        vector [][]int

        // biggest sum
        max int
    )

    // create the vector
    vector = make( [][]int, 6 )
    for i := range vector {
        vector[i] = make([]int, 6)
    }

    // read the input
    for i := 0; i < 6; i++ {
        for j := 0; j < 6; j++ {
            fmt.Scan( &vector[i][j] )
        }
    }

    // MinInt32
    max = -1 << 31

    // do the sums, and find the max
    for i := 0; i <= 3; i++ {
        for j := 0; j <= 3; j++ {

            sum := sumHourGlass(vector, i, j )
            if( sum > max ){
              max = sum
            }
        }
    }

    fmt.Println( max )
}
