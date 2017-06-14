package main
import "fmt"
import "math"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var sum_p, sum_s, aux, N int
    
    fmt.Scan(&N);
    
    for i := 0; i < N; i++{
        for j := 0; j < N; j++{
            
            fmt.Scan(&aux);
            
            if( i == j ){
                sum_p += aux;
            }
            
            if( i+j == N-1 ){
                sum_s += aux;
            }
        }
    }
    
    fmt.Print( math.Abs( (float64)(sum_p - sum_s) ) );
}
