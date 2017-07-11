def median( elements ):

    numElements = len( elements )

    if( numElements % 2 == 0 ):
        return ( ( elements[numElements//2] + elements[numElements//2-1] ) / 2 )
    else:
        return ( elements[numElements//2] )
# end median

if __name__ == '__main__':

    # size of the array
    size = int( input() )
    # array of elements
    array = list( map(int,input().split()) )

    # sort the array
    array.sort()

    med = median( array )

    lowerHalf = [ x for x in array if x < med  ]
    upperHalf = [ x for x in array if x > med  ]

    print( int( median( lowerHalf ) ) )
    print( int( med ) )
    print( int( median( upperHalf ) ) )
