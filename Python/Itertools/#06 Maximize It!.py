from itertools import combinations
from itertools import permutations
from itertools import product

if __name__ == '__main__':
    k, m = map( int, input().split() )

    listas = []

    for _ in range(k):

        l = list( map( int, input().split() ) )

        # pow all elements from this line
        l = [ pow(x,2) for x in l ]
        #l = [ x for x in l ]

        # remove the first element form the list
        l.pop( 0 )

        listas.append( sorted(l) )

    # end for

    # all possible combinations
    combinations = product(*listas)

    lista = map(sum, product(*listas))
    lista = [ x%m for x in lista ]

    print( max( list(lista) ) )
