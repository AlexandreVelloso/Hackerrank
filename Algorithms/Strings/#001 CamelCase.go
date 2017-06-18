package main
import "fmt"

func main() {
    var (
        str string
        count int
    )
    
    count = 0
    
    fmt.Scan( &str )
    
    // count the number of letters uppercase
    for i:= 0; i < len(str); i++ {
        if( str[i] >='A' && str[i] <='Z' ){
            count++;
        }
    }
    
    fmt.Println( count + 1)
}
