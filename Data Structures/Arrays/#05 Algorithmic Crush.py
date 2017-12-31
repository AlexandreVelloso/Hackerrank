if __name__ == '__main__':

    # size  -> size of the array
    # numOp -> numer of operations
    size, numOp = map( int, input().split() )

    array = []

    # initialize the array with 0's
    for _ in range( size ):
        array.append( 0 )
    # end for

    intervalos = list()

    # for each operation
    for _ in range( numOp ):
        a,b,k = map( int, input().split() )

        intervalos.append( (a,b,k) )

        for x in range( a-1, b ):
            array[x] += k
    # end for

    print( eval("1 in ({},{})".format( 1,2 ) ) )
    print( intervalos )

    print( max(array) )

    # for each operation

#      | 1   2   3   4   5
# (1,2) 100 100   0   0   0
# (2,5)   0 100 100 100 100
# (3,4)   0   0 100 100   0
