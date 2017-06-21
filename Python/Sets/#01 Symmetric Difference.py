if __name__ == '__main__':

    # i don't need this line
    input()
    M = set( input().split() )
    # i don't need this line too
    input()
    N = set( input().split() )

    # difference between the sets, this function make
    # a.difference(b) and b.difference(a).
    # and M^N also works
    difference = set( M.symmetric_difference(N) )

    # transform the set to a list of integers
    difference = list( map(int, difference) )

    # print the list sorted
    print( *sorted(difference) , sep = '\n' )
