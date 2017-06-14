package main
import "fmt"

func factorial( n int ) int{

    var fact int

    if( n <= 1 ){
        fact = 1
    }else{
        fact = n * factorial( n-1 )
    }

    return fact
}

func main() {
    var n int

    fmt.Scan( &n )

    fmt.Println( factorial(n) )

}
