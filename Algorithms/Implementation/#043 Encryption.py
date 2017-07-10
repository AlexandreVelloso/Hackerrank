import math

if __name__ == '__main__':

    # sentence to be encrypted
    sentence = input().replace(' ','')

    # number of rows
    rows = math.ceil( math.sqrt( len(sentence) ) )
    # number of columns
    columns = math.ceil( math.sqrt( len(sentence) ) )

    count = 0
    # grid to encrypt
    grid = []

    # make the grid
    for _ in range( rows ):
        grid.append( sentence[ count : columns+count ] )
        count += columns
    # end for

    encrypted = ""

    # encrypt the sentence
    for col in range(columns):
        for row in range( rows ):

            # sometimes the column don't exist
            try:
                encrypted += grid[ row ][ col ]
            except( Exception ):
                pass
        encrypted += " "

    print( encrypted.strip() )
