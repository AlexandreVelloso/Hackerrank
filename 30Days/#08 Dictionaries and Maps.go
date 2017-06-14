package main
import (
    "fmt"
)

func main() {
    var(
        size int

        name string
        number int32

        exist bool

        // map string to int
        dictionary map[string]int32
    )

    fmt.Scan( &size )

    dictionary = make( map[string]int32 )

    // fill the dictionary
    for i:= 0; i < size ; i++ {

        fmt.Scan( &name )
        fmt.Scan( &number )

        dictionary[name] = number
    }

    // search by name
    for i := 0; i < size ; i++ {

        // read the search name
        fmt.Scan( &name )

        number, exist = dictionary[name]

        // if the name exists
        if( exist ){
            fmt.Printf( "%s=%d\n",name, number )
        }else{
            fmt.Println("Not found")
        }
    }
}