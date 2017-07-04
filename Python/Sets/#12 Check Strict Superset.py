if __name__ == '__main__':

    # set A
    A = set( map(int, input().split() ) )
    # number of other sets
    num = int( input() )
    #
    resp = list()

    for _ in range(num):
        otherSet = set( map(int, input().split() ) )
        resp.append( otherSet <= A and len(otherSet) < len(A) )
    # end for

    # verify if all resp is true
    print( all(resp) )
