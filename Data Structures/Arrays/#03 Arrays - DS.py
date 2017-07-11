if __name__ == '__main__':

    # size   -> size of the array
    # numRot -> num of rotations
    size, numRot = map( int, input().split() )

    # array
    array = list( map(int, input().split()) )

    # position of the first element
    x = ( numRot )%size

    # array rotated
    rotated = []

    for _ in range( size ):

        # append element in the array
        rotated.append( array[x] )
        # pass to the next element
        x = ( x + 1 ) % size
    # end for

    # print rotated array
    print( *rotated )
