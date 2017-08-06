import numpy as np

# read the array from input
array = np.array( input().split(), float )

print( np.floor( array ) )
print( np.ceil( array ) )
print( np.rint( array ) )
