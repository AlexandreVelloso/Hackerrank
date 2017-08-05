import numpy as np

# convert the list to a numpy array
matrix = np.array( input().split(), int )

# change the shape
matrix.shape = (3,3)

print( matrix )
