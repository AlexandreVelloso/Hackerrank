from itertools import groupby

word = input()

print( *[ ( len(list(lista)), int(element) ) for element,lista in groupby(word) ] )
