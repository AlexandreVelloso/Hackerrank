package main
import (
  "fmt"
)

func main() {
    var(
        text string
        heights map[uint8]int

        h int
        caractere uint8

        max int
    )

    caractere = 97
    heights = make( map[uint8]int )
    // read the letters heights
    for i := 0; i < 26; i++{
        fmt.Scan( &h )

        heights[caractere] = h
        caractere++
    }

    // read the text
    fmt.Scan( &text )

    // height of the biggest word
    max = 0
    for i := range text{

        if( heights[ text[i] ] > max ){
            max = heights[ text[i] ]
        }

    }

    // calculate the square
    fmt.Println( max*len(text) )
}
