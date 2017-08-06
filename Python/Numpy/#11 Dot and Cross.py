import numpy as np

# dimension os the arrays, they are NxN both
n = int( input() )

# read the A array
A = np.array( [ input().split() for _ in range (n) ], int )
# read the B array
B = np.array( [ input().split() for _ in range (n) ], int )

print( np.dot(A,B) )
