package main
import "fmt"

func main() {
    var(
        num_elementos int
    )
    
    fmt.Scan( &num_elementos )
    
    for i := 1 ; i <= num_elementos ; i++ {
        
        for j := 0 ; j < num_elementos - i ; j++ {
            
            fmt.Print(" ")
            
        }
        
        for j := 0 ; j < i ; j++ {
            fmt.Print("#")
        }
        
        fmt.Println("")
        
    }
}
