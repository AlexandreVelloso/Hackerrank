package main
import "fmt"

func main() {
    // entregue antes da data -> 0
    // entregue depois do dia mas no mesmo mes e ano -> 15*num dias
    // entregue depois do mes, mas no mesmo ano -> 500 * numero de meses
    // senao 10,000
    var(
      // date expected
      day_exp, month_exp, year_exp int;
      // date actually
      day_act, month_act, year_act int;

      fine int;
    )

    // read the input
    fmt.Scan( &day_act, &month_act, &year_act);
    fmt.Scan( &day_exp, &month_exp, &year_exp );

    if( year_act < year_exp ){
      fine = 0;
    }else if( year_act == year_exp ){

      if( month_act < month_exp ){
        fine = 0;
      }else if( month_act == month_exp ){

        if( day_act <= day_exp ){
          fine = 0;
        }else{
          fine = 15*( day_act - day_exp );
        }// end if day

      }else{
        fine = 500*( month_act - month_exp );
      }// end if month
    }else{
      fine = 10000;
    }// end if year

    fmt.Println( fine );
}
