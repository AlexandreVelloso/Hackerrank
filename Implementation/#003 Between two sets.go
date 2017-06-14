package main

import "fmt"

// count the number of elements in array that are factor of x
func factorA ( x int, vet [10]int ) (int) {

    count := 0

    for i := 0; i < 10; i++ {
        if( vet[i] == 0  ){
            break
        }else if( x % vet[i] == 0 ){
            count++
        }
    }

    return count
}

// count the number of elements that have x as factor
func factorB( x int, vet[10]int ) (int){

    count := 0

    for i := 0; i < 10; i++ {
        if( vet[i] == 0 ){
            break
        }else if( vet[i] % x == 0 ){
            count++
        }
    }

    return count
}

// find the biggest element in a array
func max( vet[10]int ) (int){
    max := vet[0]

    // find the biggest value in vectorA
    for i := 1; i < 10; i++ {
        if( vet[i] > max ){
            max = vet[i]
        }
    }

    return max
}

func main() {

    var(
        vetA,vetB [10]int
        sizeA, sizeB int
    )

    // read the size of vector A and vectorB
    fmt.Scanf("%d %d",&sizeA,&sizeB)

    // read the elements of vectorA
    for i := 0; i < sizeA; i++ {
        fmt.Scan(&vetA[i])
    }

    // read the elements of vectorB
    for i := 0; i < sizeB; i++ {
        fmt.Scan(&vetB[i])
    }

    // biggest element in vectorA
    maxA := max(vetA)
    // biggest element in vectorB
    maxB := max(vetB)


    countNumbers := 0

    for i := maxA; i <= maxB; i++ {

        if( factorA(i,vetA) == sizeA && factorB(i,vetB) == sizeB ){
            countNumbers++
        }

    }

    fmt.Println(countNumbers)

}