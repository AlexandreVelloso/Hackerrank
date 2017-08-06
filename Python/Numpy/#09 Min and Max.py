import numpy as np

# n -> number of lines of the array
# m -> number of columns of the array
n, m = map( int, input().split() )

# read the array
array = np.array( [ input().split() for _ in range( n ) ], int )

# make the sum of the array
arrayMin = np.min( array, axis = 1 )

# print the product
print( np.max( arrayMin ) )
