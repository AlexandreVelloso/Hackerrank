if __name__ == '__main__':

    letras = input()
    # array para guardar o numero de ocorrencias de cada caractere
    occ = [ ]

    # cria um vetor com um caractere e a sua ocorrencia
    for x in range( 26 ):
        occ.append( [chr(x+97), 0] )

    # conta o numero de ocorrencias de cada caractere
    for c in letras:
        occ[ ord(c) - 97 ][1] += 1

    # ordena os caracteres pela ocorrencia
    # se a ocorrencia for a mesma, os caracteres
    # vao acabar ficando ordenados por ordem alfabetica
    # por que o vetor occ ja esta em ordem alfabetica
    occ.sort( key = lambda x: x[1], reverse=True )

    # mostra as 3 letras mais comuns
    print( *[ "{} {}".format( occ[x][0], occ[x][1] ) for x in range(3) ], sep = '\n' )


'''
Using collections

Author: Boki

from collections import Counter, OrderedDict

class OrderedCounter(Counter, OrderedDict):
    pass
[print(*c) for c in OrderedCounter(sorted(input())).most_common(3)]
'''
