from collections import OrderedDict

if __name__ == '__main__':

    ordered_dictionary = OrderedDict()

    # for N items
    for _ in range( int(input()) ):
        splited = input().strip().split()

        # the item name maybe have spaces,
        # i need to concat all name
        #
        # the join concat all list to a string
        item_name = ' '.join( splited[:-1] )

        # net price of the item
        net_price = int( splited[-1])

        # if the item_name exists in the dictionary
        if( item_name in ordered_dictionary.keys() ):
            # update the net_price
            ordered_dictionary[ item_name ] += net_price
        else:
            # insert the item in the dictionary
            ordered_dictionary [ item_name ] = net_price

    # end for

    print( *[ "{} {}".format( x[0], x[1] ) for x in ordered_dictionary.items() ], sep = '\n' )
