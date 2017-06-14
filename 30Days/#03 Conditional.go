package main
import "fmt"

func main() {
    var number int8

    fmt.Scan(&number)

    if( number % 2 == 1){
        fmt.Println("Weird")
    }else if( number >=2 && number <= 5){
        fmt.Println("Not Weird")
    }else if( number >= 6 && number <= 20 ){
        fmt.Println("Weird")
    }else if( number > 20 ){
        fmt.Println("Not Weird")
    }
}