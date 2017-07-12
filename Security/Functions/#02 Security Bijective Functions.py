from collections import Counter

if __name__ == '__main__':

    # number of elements
    n = int( input() )

    # array of elements
    array = list( map(int, input().split() ) )

    # count the number of occurrences of all numbers
    count = Counter( array )

    # value to be printed in output
    dictionary = { False: "YES", True:"NO" }

    print( dictionary[ any ( [ (x > 1) for x in count.values() ] ) ] )
