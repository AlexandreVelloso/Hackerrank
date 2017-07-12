if __name__ == '__main__':

    # number of elements
    numElements = int( input() )
    # array of elements
    calories = list( map( int, input().split() ) )

    # to minimize the miles, Marc have to eat
    # the most caloric cupcakes first
    calories.sort( reverse = True )

    numCupcakesEaten = 0
    miles = 0

    for x in range( numElements ):
        miles = miles + ( calories[x] * (2**numCupcakesEaten) )
        numCupcakesEaten += 1
    # end for

    print( miles )
