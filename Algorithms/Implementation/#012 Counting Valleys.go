package main
import "fmt"

func main() {
    var (
        sea_level int
        count_valley int
        num_step int
        steps string
    )

    // sea level is in 0
    sea_level = 0
    // number of valleys
    count_valley = 0

    fmt.Scan( &num_step )
    fmt.Scan( &steps )

    for i:= 0; i < num_step; i++ {
        if( steps[i] == 'D' ){

           // if sea level was 0, then it is a valley
           if( sea_level == 0 ){
               count_valley ++
           }

           sea_level --
        }else{

            sea_level ++

        }
    }

    fmt.Println( count_valley )

}
