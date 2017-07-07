if __name__ == '__main__':

    # number of test cases
    numTest = int( input() )

    for _ in range( numTest ):

        # number of black and white gifts
        numBlack, numWwhite = map( int, input().split() )

        # costs
        costBlack, costWhite, costChange = map( int, input().split() )

        # minimum price of a white gift
        costBlack = min( costBlack, costWhite + costChange )

        # minimum price of a white gift
        costWhite = min( costWhite, costBlack + costChange )

        print( (numBlack * costBlack) + (numWwhite * costWhite) )

    # end for

# end main
