package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT

    var(
        num int
    )

    // read the number
    fmt.Scan( &num )

    // show the multiples
    for i := 1 ; i <= 10; i++{

        fmt.Printf("%d x %d = %d\n",num,i, (num*i) )

    }
}