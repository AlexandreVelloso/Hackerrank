def median( elements ):

    numElements = len( elements )

    if( numElements % 2 == 0 ):
        return ( ( elements[numElements//2] + elements[numElements//2-1] ) / 2 )
    else:
        return ( elements[numElements//2] )
# end median

if __name__ == '__main__':
    # number of different elements
    numElements = int( input() )

    # array of elements
    elements = list( map(int,input().split()) )
    # array of frequences
    frequences = list( map(int,input().split()) )

    array = []

    # dataset S in the example, i will name it array
    for x in range( numElements ):
        for _ in range( frequences[x] ):
            array.append( elements[x] )
    # end for

    array.sort()

    lowerHalf = array[ 0: len(array)//2 ]

    if( len(array) % 2 == 0 ):
        upperHalf = array[ len(array)//2 : ]
    else:
        upperHalf = array[ len(array)//2 + 1 : ]

    medianLower = median( lowerHalf )
    medianUpper = median( upperHalf )

    print( "%.1f" % (medianUpper - medianLower) )
