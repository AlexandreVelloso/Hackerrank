if __name__ == '__main__':
    N,M = map( int, input().split() )

    elements = [ input().split() for _ in range( N ) ]

    K = int( input() )

    elements.sort( key=lambda x: int( x[K] ) )

    for e in elements:
        print( *e )
