import numpy as np

# n -> number of lines
# m -> number of columns
n, m = map( int, input().split() )

# read the input as a numpy array
array = np.array( [ input().split() for _ in range( n ) ], int )

print( np.transpose( array ) )
print( array.flatten() )
