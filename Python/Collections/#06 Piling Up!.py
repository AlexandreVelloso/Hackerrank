from collections import deque

if __name__ == '__main__':

    for _ in range( int(input()) ):

        d = deque()

        # sideLength of the cube
        sideLength = int( input() )

        lista = list( map( int, input().split() ) )

        for x in range( sideLength ):
            d.append( lista[x] )

        while( len(d) > 0 ):

            r = d.pop()
            l = d.popleft()

            if( r > l ):
                break
        # end while

        if( len(d) == 0 ):
            print( "Yes" )
        else:
            print( "No")

'''
Nao entendi o problema
'''
