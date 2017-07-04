package main
import (
    "fmt"
    "bufio"
    "os"
    "strconv"
)

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var line string
    
    reader := bufio.NewReader(os.Stdin)
    
    // readline as string
    line, _ = reader.ReadString('\n')
    
    // read the hour and convert to int
    hour, _ := strconv.ParseInt( string( line[0] ) + string( line[1] ), 10, 32 )
    
    // read AM or PM
    time := string( line[8:10] )
    
    // if the hour is PM, convert
    if( time == "AM" ){
        if( hour == 12 ){
            hour = 0
        }
    }else if( hour != 12 ){
        hour = (hour + 12)%24
    }
    
    // get the minutes and seconds
    resto := string( line[3:8])
    
    if( hour < 10 ){
     fmt.Printf("0%d:%s\n", hour, resto )
    }else{
        fmt.Printf("%d:%s\n", hour, resto)
    }
    
}
