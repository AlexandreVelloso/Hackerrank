package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var A [3] int
    var score_A, score_B int
    var aux int
    
    for i := 0; i < 3; i++{
        fmt.Scanf("%d", &A[i])
    }
    
    for i := 0; i < 3; i++ {
        fmt.Scanf("%d",&aux)
        if( A[i] > aux ){
            score_A ++;
        }else if( A[i] < aux ){
            score_B ++;
        }
    }
    
    fmt.Printf("%d %d", score_A, score_B);
    
}
