if __name__ == '__main__':

    # number of elements
    numElements = int( input() )
    # array of elements
    array = []
    # max distance
    MAX = 10**4

    x = 0
    # all elements are stored by ( number, position )
    for el in list( map( int, input().split() ) ):
        array.append( (el,x) )
        x += 1
    # end for

    # sort the array by element, if the element is equal
    # sort by positon
    array.sort( key = lambda x: (x[0], x[1]) )

    minDistance = MAX

    for x in range( numElements-1 ):

        # if they are the same number
        if( array[x][0] == array[x+1][0] ):
            minDistance = min( minDistance, abs( array[x][1] - array[x+1][1] ) )
        # end else

    # end for

    if( minDistance == MAX ):
        print("-1")
    else:
        print( minDistance )
