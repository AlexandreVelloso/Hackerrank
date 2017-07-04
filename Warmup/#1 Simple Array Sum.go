package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var size, aux, sum int
    fmt.Scanf("%d\n",&size)
    
    for i := 0; i < size; i++ {
        fmt.Scanf("%d",&aux)
        sum += aux
    }
    
    fmt.Printf("%d",sum)
}
