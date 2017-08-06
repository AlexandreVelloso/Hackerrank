import numpy as np

tupla = tuple( map( int, input().split() ) )

# comparing with the input, i need to print this
# print( np.zeros( (3,3,3), np.int) )
# print( np.ones( (3,3,3), np.int) )

# then, the solution is
print( np.zeros( tupla, np.int) )
print( np.ones( tupla, np.int) )
