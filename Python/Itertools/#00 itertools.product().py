from itertools import product

A, B = ( list(input().split()) for _ in range(2) )

print( *[ "({0}, {1})".format( x[0], x[1] ) for x in list( product(A,B) ) ], sep = ' ' )
