package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var(
        max int32
        size int
        num int32
    )
    
    // number of elements
    fmt.Scan(&size)
    
    array := make( []int32, size )
    
    // initialize the max with the first element
    fmt.Scan(&num)
    max = num
    array[0] = num
    
    // read elements and find the highest height candle
    for i:= 1; i < size; i++ {
        
        fmt.Scan(&num)
        array[i] = num
        if( num > max ){
            max = num
        }
    }
    
    // cont the numbers of occurrences of max
    cont := 0
    
    // count the number of occurrences of the max in the array
    for i := 0; i < size; i++ {
        if( array[i] == max ){
            cont++
        }
    }
    
    fmt.Println(cont)
}
