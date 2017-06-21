if __name__ == '__main__':

    # i don't need this line
    input()

    # set A
    A = set( map( int, input().split() ) )

    # number of operations
    numOp = int( input() )

    for _ in range( numOp ):

        # read the operation from the input
        operation = input().split()

        # read the set from the input
        otherSet = set( map( int, input().split() ) )

        # do the operation
        eval('A.{0}(otherSet)'.format( operation[0] ) )

    print( sum(A) )

'''
.update() or |=
.intersection_update() or &=
.difference_update() or -=
.symmetric_difference_update() or ^=
'''
