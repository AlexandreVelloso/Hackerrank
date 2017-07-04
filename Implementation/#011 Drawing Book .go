package main
import "fmt"

func main() {
   var ( n_pages, page int )

    fmt.Scan( &n_pages )
    fmt.Scan( &page )

    // number of pages to turn starting in the first page
    // this is the last page
    n := n_pages /2

    // this is the page that teacher asked
    p := page / 2

    // number of pages to turn stating in the last page
    delta := n-p

    // if the page 'p' is the same of the last page
    if( delta == 0 ){
        fmt.Println(0)
    }else if( delta > p ){

        // if is best to start from the begin
        fmt.Println( p )
    }else{

        // if is best to start from the end
        fmt.Println( delta )
    }
}

/*

page page
  0   1     2x0 2x0+1
  2   3     2x1 2x1+1
  4   5     2x2 2x2+1
  6         2x3

e.g:
    n_pages = 6
    page = 3

    n = 6/2 = 3

    p = 3/2 = 1

    // number of pages to turn stating in the last page
    delta = n - p = 3 - 1 = 2

    is best to start from the begin, because p < delta
*/