#
def nextMove(n,r,c,grid):

    # position of the bot in the grid
    posM = (r,c)
    # initialize the position of the princess in the grid
    posP = (-1,-1)
    # result to be printed
    result = ""

    # find the position of the princess( p )
    for x in range( n ):
        if( 'p' in grid[x]):

            for y in range(n):
                if( grid[x][y] == 'p'):
                    posP = (x,y)
    # end for

    if( posM[0] == posP[0] ):
        if( posM[1] > posP[1] ):
            result = "LEFT"
        else:
            result = "RIGHT"

    elif( posM[1] == posP[1] ):
        if( posM[0] > posP[1] ):
            result = "UP"
        else:
            result = "DOWN"
    else:
        # posM != posP
        if( posM[0] < posP[0] ):
            result = "DOWN"
        else:
            result = "UP"

    return result
# end nextMove

n = int(input())
r,c = [int(i) for i in input().strip().split()]
grid = []
for i in range(0, n):
    grid.append(input())

print(nextMove(n,r,c,grid))
