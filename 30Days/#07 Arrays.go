package main
import "fmt"

func main() {
    var(
        // size of the array
        size int
    )

    fmt.Scan( &size )

    array := make( []int, size)

    // fill the array
    for i:=0; i < size; i++ {
        fmt.Scan( &array[i] )
    }

    // show the elements of the array in a reverse order
    for i := size-1; i >= 0; i--{
        fmt.Printf("%d ",array[i])
    }
    fmt.Println()

}