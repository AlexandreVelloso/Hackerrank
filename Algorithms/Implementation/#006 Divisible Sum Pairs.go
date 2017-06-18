package main
import "fmt"

func main() {
    var(
        // size of the array
        size int
        // integer k
        k int
    )

    fmt.Scan(&size)
    fmt.Scan(&k)

    // number of divisibles by k
    count_div := 0

    vet := make( []int, size)

    // read the vector elements
    for i := 0; i < size; i++ {
        fmt.Scan( &vet[i] )
    }

    // pairs of i and j
    for i := 0; i < size; i++{
        for j := i+1; j < size; j++ {

            // if Ai + Aj is divisible by k
            if( ( vet[i] + vet[j] )% k == 0 ){
                count_div++
            }
        }
    }

    fmt.Println( count_div )
}