package main
import "fmt"

func main() {
 
    var(
        house_s int
        house_t int
        
        apple_tree int
        orange_tree int
        
        number_apples int
        number_oranges int
        
        apple int
        orange int
    )
    
    fmt.Scan( &house_s )
    fmt.Scan( &house_t )
    
    fmt.Scan( &apple_tree )
    fmt.Scan( &orange_tree )
    
    fmt.Scan( &number_apples )
    fmt.Scan( &number_oranges )
    
    num_apple_house := 0
    num_orange_house := 0
    
    // count the number of apples that fall in
    // the house
    for i := 0; i < number_apples ;i++ {
        fmt.Scan( &apple )
        
        fall := (apple_tree + apple )
        
        if( fall >= house_s && fall <= house_t ){
            num_apple_house ++
        }
    }
    
    // count the number of oranges that fall in
    // the house
    for i := 0; i < number_oranges ;i++ {
        fmt.Scan( &orange )
        
        fall := (orange_tree + orange )
        
        if( fall >= house_s && fall <= house_t ){
            num_orange_house ++
        }
    }
    
    fmt.Println(num_apple_house)
    fmt.Println(num_orange_house)
    
}
