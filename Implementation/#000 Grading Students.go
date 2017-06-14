package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    
    var(
        size int
        num int
        
    )
    
    fmt.Scan( &size )
    
    for i:= 0; i< size; i++ {
        fmt.Scan( &num )
        
        // round the number
        if( num >= 36 ){
            
            mod := num%5
            
            // if mod <= 2, the closest multiple of 5 is
            // lower than the num
            if( mod > 2){
                // round the number
                num += (5 - mod)
            }
            
        }
        
        fmt.Println(num)
    } 
}
