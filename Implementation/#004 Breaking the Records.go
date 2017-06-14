package main
import "fmt"

func main() {
    var(
        n, num, max, min, cMax, cMin int
    )

    // read the number of games
    fmt.Scan( &n )

    // read the first game, it will be the first
    // record
    fmt.Scan( &num )

    min = num
    max = num

    cMax = 0
    cMin = 0

    for i := 1; i < n; i++{
        fmt.Scan( &num )

        if( num > max ){
            cMax++
            max = num
        }

        if( num < min ){
            cMin++
            min = num
        }
    }

    fmt.Printf("%d %d",cMin, count_least)
}