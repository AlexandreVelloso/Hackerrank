if __name__ == '__main__':

    # dollar = n = dollars
    # costChoc = c = chocolate cost
    # costChange = m = chocolate wrappers in exchange
    # numChoc = numer of chocolates that Booby have

    # for each test case
    for _ in range( int(input()) ):
        dollar, costChoc, costChange = map( int, input().split() )

        # buy how many chocolates Bobby can
        numChoc = dollar // costChoc
        # each chocolate give a wrap
        numWrap = numChoc

        while( numWrap >= costChange ):
            numWrap = numWrap - costChange + 1
            numChoc += 1
        # end while

        print( numChoc )
