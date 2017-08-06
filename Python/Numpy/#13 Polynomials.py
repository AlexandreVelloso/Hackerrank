import numpy as np

# read the coeficients
coeficients = list( map( float, input().split() ) )
# read the x value
x = int( input() )

print( np.polyval( coeficients, x ) )
