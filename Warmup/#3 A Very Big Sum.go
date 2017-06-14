package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var sum uint64
    var num_elements int
    var aux uint64
    
    fmt.Scanf("%d",&num_elements);
    sum = 0;
    
    for i := 0; i < num_elements; i++ {
        fmt.Scanf("%d", &aux);
        sum += aux;
    } 
    
    fmt.Printf("%d",sum)
}
