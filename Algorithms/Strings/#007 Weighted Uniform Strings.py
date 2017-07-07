# this method make a list of each uniform string
# Example:
# "abccdddeaa"
# a occurs 1 time ( a,1 )
# b occurs 1 time( b,1 )
# c occurs 2 times( c,2 )
# d occurs 3 times( d,2 )
# e occurs 2 times( e,2 )
# a occurs 2 times( a,2 )
def makeList ( word ):
    lista = list()
    count = 1

    for pos in range( len(word) ):
        char = word[pos]
        nextC = word[pos+1]

        if( pos < ( len(word) -1 ) and nextC == char ):
            pos += 1
            count +=1
        else:
            lista.append( (char,count) )
            count = 1

    # end for

    return lista
# end method

if __name__ == '__main__':

    count = makeList( input() )

    # number of test cases
    numberCases = int( input() )

    # for each test case
    for _ in range( numberCases ):

        # wheight from input to be tested
        weight = int( input() )

        found = False

        # for each tuple( character, occurrences )
        for c in count:
            # weight of the character
            num = ord( c[0] ) - 96

            if( weight % num == 0 and weight <= ( num * c[1] ) ):
                print("Yes")
                found = True
                break
        # end for

        if( not found ):
            print( "No" )

    # end for
