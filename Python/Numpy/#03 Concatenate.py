import numpy as np

# n -> number of lines of the first array
# m -> number of lines of the second array
# p -> number of columns of the two arrays
n, m, p = map( int, input().split() )

# read the first array
array1 = np.array( [ input().split() for _ in range (n) ], int )
# read the second array
array2 = np.array( [ input().split() for _ in range (m) ], int )

# concatenate the arrays
print( np.concatenate( (array1 , array2), axis = 0 ) )
