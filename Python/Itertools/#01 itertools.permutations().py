from itertools import permutations

word, k = input().split()
k = int( k )

lista = ["".join( x ) for x in permutations( sorted(word) ,k) ]

print( *lista, sep = '\n' )
