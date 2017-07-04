package main
import(
    "fmt"
)
    
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var(
        num_elementos int
        elemento int
        
        num_positivos int
        num_zeros int
        num_negativos int
    )
    
    num_positivos = 0
    num_zeros = 0
    num_negativos = 0
    
    fmt.Scan( &num_elementos )
    
    for i := 0 ; i < num_elementos ; i++ {
        fmt.Scan( &elemento )
        
        if( elemento > 0 ){
            num_positivos ++;
        }else if( elemento == 0 ){
            num_zeros ++;
        }else{
            num_negativos ++;
        }
    }
    
    
    fmt.Printf("%.6f\n", float64(num_positivos)/float64(num_elementos) )
    fmt.Printf("%.6f\n", float64(num_negativos)/float64(num_elementos) )
    fmt.Printf("%.6f\n", float64(num_zeros)/float64(num_elementos) )
}
