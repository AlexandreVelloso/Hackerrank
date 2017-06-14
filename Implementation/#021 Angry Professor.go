package main
import "fmt"

func main() {
  var(
    // number of test cases
    num_cases int

    // number of students in the class
    num_students int
    // cancelation threshold
    cancel int

    // time of a student
    time int
  )

  fmt.Scan( &num_cases )

  for i:=0 ; i < num_cases; i++ {
    fmt.Scan( &num_students )
    fmt.Scan( &cancel )

    // number of students in the class
    num_ok := 0
    // read the tiem of the students, and calculate
    // how many students are late
    for x := 0; x < num_students; x++ {

      fmt.Scan( &time )

      // if the student is late
      if( time <= 0 ){
        num_ok++
      }

    }

    if( cancel > num_ok ){
      fmt.Println( "YES" )
    }else{
      fmt.Println( "NO" )
    }

  }
}
