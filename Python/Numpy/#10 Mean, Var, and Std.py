import numpy as np

# n -> number of lines of the array
# m -> number of columns of the array
n, m = map( int, input().split() )

# read the array
array = np.array( [ input().split() for _ in range( n ) ], int )

print( np.mean( array, axis = 1) )
print( np.var( array, axis = 0) )
print( np.std( array) )
