package main
import "fmt"

func main() {
    var(
        size int
        k int
        item_cost int
        charged int
    )
    
    // amount of items
    fmt.Scan( &size )
    // item that Anna don't eat
    fmt.Scan( &k )
    
    // read the cost of the shared items
    cost := 0
    for i := 0; i < size ; i++ {
        
        fmt.Scan( &item_cost )
        
        // if Anna didn't eat the item
        if( i != k ){
            cost += item_cost
        }
        
    }
    
    // divide the cost by 2
    cost = cost/2
    
    fmt.Scan( &charged )
    
    r := charged - cost
    
    if( r == 0 ){
        fmt.Println( "Bon Appetit" )
    }else{
        fmt.Println( r )
    }
}
