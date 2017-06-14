package main
import (
    "fmt"
    "bufio"
    "os"
)

func main(){
    fmt.Printf("Hello World!\n")
    reader := bufio.NewReader(os.Stdin)
    line, _ := reader.ReadString('\n')
    fmt.Println(line)
}