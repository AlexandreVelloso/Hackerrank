#!/usr/bin/python

def printMove( move, n ):
    for x in range( n ):
        print( move )

def displayPathtoPrincess(n,grid):
    posM = (-1,-1)
    posP = (-1,-1)

    # find the position of the bot( m )
    for x in range( n ):
        if( 'm' in grid[x]):

            for x in range(n):
                if( grid[x][y] == 'm'):
                    posM = (x,y)
    # end for

    # the princess is in one corner,
    # find the position of the princess
    if( grid[0][0] == 'p' ):
        posP = (0,0)
    elif( grid[n-1][0] == 'p' ):
        posP = (n-1,0)
    elif( grid[0][n-1] == 'p' ):
        posP = (0, n-1)
    elif( grid[n-1][n-1] == 'p' ):
        posP = (n-1, n-1)

    #*****************************

    if( posM[0] > posP[0] ):
        printMove( "UP", (posM[0] - posP[0]))
    elif( posM[0] < posP[0] ):
        printMove( "DOWN", (posP[0] - posM[0]) )

    if( posM[1] > posP[1] ):
        printMove( "LEFT", (posM[1] - posP[1] ) )
    elif( posM[1] < posP[1] ):
        printMove( "RIGHT", (posP[1] - posM[1] ) )

# end method

m = int(input())
grid = []
for i in range(0, m):
    grid.append(input().strip())

displayPathtoPrincess(m,grid)
