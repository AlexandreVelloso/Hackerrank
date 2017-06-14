package main
import (
    "fmt"
    "bufio"
    "os"
)

func main() {

    var(
        line string
        num_cases int
        length int
    )

    reader := bufio.NewReader(os.Stdin)

    // number of cases
    fmt.Scanf("%d\n",&num_cases)

    for x := 0; x < num_cases; x++{

        // read the line
        line, _ = reader.ReadString('\n')

        // length of the string, ignore the \n
        length = len(line)

        // len(line)
        for i:= 0; i < length; i+=2{
            if( string(line[i] ) != string('\n') ){
                fmt.Print( string( line[i] ) )
            }
        }

        fmt.Print(" ")

        for i:= 1; i < length; i+=2{
            if( string(line[i] ) != string('\n') ){
                fmt.Print( string( line[i] ) )
            }
        }

        fmt.Println("")
    }

}