package main
import "fmt"

func main() {
    var(
        // size of the chocolate bar
        size int

        day,month int
    )

    // read the size of the bar
    fmt.Scan( &size )

    // create the chocolate bar
    bar := make( []int, size )

    // read the numbers
    for i :=0 ;i < size; i++ {
        fmt.Scan( &bar[i] )
    }

    // day
    fmt.Scan(&day)
    // month
    fmt.Scan(&month)

    count := 0

    // m is the number of squares that Lily wants to give to Ron
    for i := 0 ; i <= size - month ; i++ {

        sum := 0

        // sum 'month' square(s) of chocolate
        for j:= 0 ; j < month; j++ {
           sum += bar[i+j]
        }

        if( sum == day ){
            count ++
        }

    }

    fmt.Println(count)
}