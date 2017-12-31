import re

# read the roman number
number  = input()

# I, V, X, L, C, D, M
pattern = "(I{4,}|V{4,}|X{4,}|L{4,}|C{4,}|D{4,}|M{4,})"

# existe precendencia dos caracteres
print( re.search(pattern,number) == None and all( c in "IVXLCDM" for c in number ) )
