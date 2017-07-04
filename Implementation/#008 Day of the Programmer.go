package main
import "fmt"

// verify if the year Julian or Gregorian calendar
func isJulian (year int ) bool{

    var is_julian bool

    if( year >= 1700 && year <= 1917 ){
        is_julian = true
    }else{
        is_julian = false
    }

    return is_julian
}

// verify if the year is a leap year
func isLeap ( year int ) bool {

    var is_leap bool

    // julian calendar
    if( isJulian(year) ){

        // *divisible by 4
        if( year%4 == 0 ){
            is_leap = true
        }else{
            is_leap = false
        }

    }else{
        // gregorian calendar

        // *divisible by 400
        // *divisible by 4
        // * not divisible by 100
        if( ( year % 4 == 0 && year % 100 != 0 ) || ( year % 400 == 0 ) ){
            is_leap = true
        }else{
            is_leap = false
        }

    }

    return is_leap
}

func main() {
    var(
        year int
        )

    // read the year
    fmt.Scan( &year )

    // number of days for each month
    array_month := []int {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}



    // in the year 1918 february has 14 days
    if( year == 1918 ){

        array_month[1] = 15

    }else if( isLeap(year) ){

        // if the year is leap, february have 29 days
        array_month[1] = 29
    }

    sum := 0
    i := 0

    for( sum + array_month[i] <= 256 ){
        sum += array_month[i]
        i++
    }

    day := 256 - sum
    month := i+1

    fmt.Printf( "%02d.%02d.%0d\n",day,month, year )
}