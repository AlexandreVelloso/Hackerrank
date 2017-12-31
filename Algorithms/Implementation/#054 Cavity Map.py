
# number of cases
n = int(input().strip())

# grid with all cases, 1 per line
grid = []

# read all cases
for _ in range(n):
    grid_t = str(input().strip())
    grid.append(grid_t)
# end for

# for each case
for case in grid:

    # case with X replaces, first append the border left
    newMap = case[0]

    # test the cavities
    for pos in range( 1, len(case)-1 ):

        #print( case[pos-1], case[pos], case[pos+1] )
        #print( case[pos-1] < case[pos] , case[pos+1] < case[pos] , newMap[pos-1] != 'X' )
        # if there is a cavity, and the prevous is not a cavity
        if( case[pos-1] < case[pos] and case[pos+1] < case[pos] and newMap[pos-1] != 'X'):
            newMap += "X"
        else:
            newMap += case[pos]
    # end for

    # append the border right
    newMap += case[-1]
    print( newMap )

# end for
