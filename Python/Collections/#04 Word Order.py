from collections import OrderedDict

if __name__ == '__main__':

    ordered_dictionary = OrderedDict()

    for _ in range( int(input()) ):

        # read a word from input
        word = input().strip()

        if( word in ordered_dictionary.keys() ):
            ordered_dictionary[ word ] += 1
        else:
            ordered_dictionary[ word ] = 1
    # end for

    # number of distinct words
    print( len( ordered_dictionary ) )
    # occurrence of each word
    print( *[ x[1] for x in ordered_dictionary.items() ] )
