package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
)

func main() {
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)

    scanner2 := bufio.NewReader(os.Stdin)

    // Declare second integer, double, and String variables.
    var (
     var_int uint64
     var_double float64
     var_string string
    )

    retardado, _ := strconv.ParseInt("1", 10, 32)
    retardado += 1
    scanner.Err()

    fmt.Scan(&var_int)
    fmt.Scan(&var_double)
    var_string, _ = scanner2.ReadString('\n')

    var_int += i
    var_double += d
    var_string = s + var_string

    fmt.Println (var_int)
    fmt.Printf("%.1f\n",var_double)
    fmt.Println(var_string)
}
