package main
import "fmt"

func main() {
    var(
        size int
        vector [101]int

        color int
        number_pairs int
    )

    fmt.Scan( &size )

    // read the vector elements, and save in the hash structure vector
    for i:= 0; i < size; i++ {
        fmt.Scan( &color )

        vector[ color ]++
    }

    // count the number of pairs
    number_pairs = 0

    for i := 0; i < 101; i++ {
        number_pairs += vector[i]/2
    }

    fmt.Println( number_pairs )
}
