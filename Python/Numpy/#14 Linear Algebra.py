import numpy as np

# size of the matrix, it is a NxN matrix
n = int( input() )
# read the matrix
matrix = np.array( [ input().split() for _ in range( n ) ], float )

# print the determinant
print( np.linalg.det( matrix ) )
