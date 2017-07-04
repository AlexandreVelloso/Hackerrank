package main
import "fmt"

func main() {
    var(
        number_input int
        catA, catB, mouseC int
    )

    fmt.Scan(&number_input)

    for i:= 0; i < number_input; i++ {
        // read the inputs
        fmt.Scanf("%d %d %d",&catA, &catB, &mouseC)

        // distance of catA and the mouse
        distA := catA - mouseC
        if( distA < 0 ){
            distA *= -1
        }

        // distance of catB and the mouse
        distB := catB - mouseC
        if( distB < 0 ){
            distB *= -1
        }

        // test which cat will catch the mouse
        if( distA < distB ){
            fmt.Println( "Cat A" )
        }else if( distB < distA ){
            fmt.Println( "Cat B" )
        }else{
            fmt.Println( "Mouse C" )
        }
    }
}
