package main

import (
    "fmt"
)

func main(){
    var(
        mealCost float64
        tipPercent int32
        taxPercent int32
    )

    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)

    tip := ( (float64)(tipPercent)/100 * mealCost )
    tax := ( (float64)(taxPercent)/100 * mealCost )

    totalCost := mealCost + tip + tax

    fmt.Printf( "The total meal cost is %d dollars.",(int)(totalCost+0.5) )

}