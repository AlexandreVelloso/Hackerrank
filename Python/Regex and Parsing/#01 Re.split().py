import re

if __name__ == '__main__':

    number = input().strip()
    lista = list( re.split("[,.]", number) )
    print( *[ x for x in lista if( x != "")], sep = '\n' )
