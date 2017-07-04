package main
import "fmt"

func main() {
    var(
        money int
        number_keyboard int
        number_usb int

        vect_keyboard []int
        vect_usb []int
    )

    fmt.Scanf( "%d %d %d",&money, &number_keyboard, &number_usb )

    vect_keyboard = make( []int, number_keyboard)
    vect_usb = make( []int, number_usb )

    // read the prices of keyboards
    for i:= 0; i < number_keyboard; i++ {
        fmt.Scan( &vect_keyboard[i] )
    }

    // read the prices of usb
    for i := 0; i < number_usb; i++ {
        fmt.Scan( &vect_usb[i] )
    }

    max_cost := -1
    for i := 0; i < number_keyboard; i++ {
        for j := 0; j < number_usb; j++ {

            combined := vect_keyboard[i] + vect_usb[j]

            // if the items price is better
            // and Monica have money to buy these two
            if( combined > max_cost && combined <= money ){
                max_cost = combined
            }

        }
    }

    fmt.Println( max_cost )
}
