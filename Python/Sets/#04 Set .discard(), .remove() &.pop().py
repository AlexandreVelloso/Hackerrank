if __name__ == '__main__':

    # i don't need this line
    input()

    # array from input
    array = set( map( int, input().split() ) )

    # number of operations
    numOp = int( input() )

    # for each operation
    for _ in range( numOp ):

        # Author: DeadMoroz
        eval('array.{0}({1})'.format( *input().split()+[''] ) )

    # end for

    print( sum( array ) )
