import numpy as np

# read the A array
A = np.array( input().split(), int )
# read the B array
B = np.array( input().split(), int )

print( np.inner( A,B ) )
print( np.outer( A,B ) )
