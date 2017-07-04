package main
import "fmt"

func main() {

    var(
        x1, x2 int
        v1, v2 int
    )
    
    // read the input
    fmt.Scanf("%d %d %d %d",&x1,&v1,&x2,&v2)
    
    if( x2 > x1 ){
        
        // if the kangaroo x2 begin after x1 and
        // x2 has more velocity than x1, then the answer is no
        if( v2 >= v1 ){
            fmt.Println("NO")
        }else{
            
            // x2 > x1
            dif_inicial := x2 - x1
            // v1 > v2
            dif_velocidade := v1 - v2
            
            // if they are in the same location
            if( dif_inicial % dif_velocidade == 0 ){
                fmt.Println("YES")
            }else{
                fmt.Println("NO")
            }
            
        }
        
    }else if( x1 > x2 ){
        
        // if the kangaroo x1 begin after x2 and
        // x1 has more velocity than x2, then the answer is no
        if( v1 >= v2 ){
            fmt.Println("NO")
        }else{
            
            // x1 > x2
            dif_inicial := x1 - x2
            // v2 > v1
            dif_velocidade := v2 - v1
            
            // if they are in the same location
            if( dif_inicial % dif_velocidade == 0 ){
                fmt.Println("YES")
            }else{
                fmt.Println("NO")
            }
            
        }
        
    }else{
        // if they begin in the same location, and have the same velocity,
        // they will land on the same location in the first jump, else they
        // will never land on the same location
        if( v1 == v2 ){
            fmt.Println("YES")
        }else{
            fmt.Println("NO")
        }
    }
    
}
