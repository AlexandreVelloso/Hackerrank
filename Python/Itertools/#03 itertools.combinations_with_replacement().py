from itertools import combinations

word, k = input().split()
k = int( k )
word = sorted( word )

print( *["".join( x ) for x in combinations( word ,k) ], sep = '\n' )
