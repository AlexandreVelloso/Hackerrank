if __name__ == '__main__':

    # i don't need this numbers
    n, m = input().split()
    # read the array from input
    array = list( map( int, input().split() ) )
    # set A
    A = set( map( int, input().split() ) )
    # set B
    B = set( map( int, input().split() ) )

    # if ( value in A ) then return 1, else return 0
    happinessA = sum( [value in A for value in array ] )
    happinessB = sum( [value in B for value in array ] )

    print( happinessA - happinessB )
