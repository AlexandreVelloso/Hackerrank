package main
import "fmt"

func main() {
    var(
        // number of birds
        size int
    )

    fmt.Scan( &size )

    birds := make( []int, size)

    // read the birds
    for i := 0; i < size; i++ {
        fmt.Scan( &birds[i] )
    }

    // number of birds of each specie
    counter := make( []int, size)

    // count the number of birds of each specie
    for i := 0; i < size; i++ {
        counter[ birds[i]-1 ]++
    }


    // id of the most common bird
    id_common := 1
    max := counter[0]

    // find the most common bird ID
    for i := 1; i < size; i++ {
        if( counter[i] > max ){
            id_common = i+1
            max = counter[i]
        }
    }

    fmt.Println( id_common )

}