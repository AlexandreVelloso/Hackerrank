import numpy as np

# n -> number of lines of the two arrays
# m -> number of columns of the two arrays
n, m = map( int, input().split() )

# read the first array
A = np.array( [ input().split() for _ in range (n) ], int )
# read the second array
B = np.array( [ input().split() for _ in range (n) ], int )

print( A + B )
print( A - B )
print( A * B )
print( A // B )
print( A % B )
print( A ** B )
