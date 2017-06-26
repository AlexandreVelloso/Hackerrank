from itertools import combinations

word, k = input().split()
k = int( k )
word = sorted( word )

for i in range(1, k+1):
    print( *["".join( x ) for x in combinations( word ,i) ], sep = '\n' )
