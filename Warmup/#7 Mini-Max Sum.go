package main
import "fmt"
import "math"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var(
        array [ 5 ] int64
        min int64
        max int64
        
        sum int64
    )
    
    // read the data
    for i := 0 ; i < 5; i++ {
        fmt.Scan( &array[i] )
    }
    
    min = math.MaxInt64
    max = math.MinInt64
    
    for i := 0 ; i < 5 ; i++ {
        
        sum = 0
       
        // sum elements, ignoring 'i'
        for j := 0 ; j < 5 ; j++ {
            
            if( i != j ){
                sum += array[j]
            }
            
        }
        
        if( sum < min ){
            min = sum
        }
        
        if( sum > max ){
            max = sum
        }
    }
    
    fmt.Printf("%d %d\n", min, max)
}
