from collections import defaultdict

if __name__ == '__main__':

    # integers N and M
    N, M = map( int, input().split() )

    groupA = defaultdict( list )
    groupB = list()

    # for each word in group A
    for x in range( N ):

        # I'th word
        word = input()

        groupA[ word ].append( x+1 )
    # end for

    for x in range( M ):
        # I'th word
        groupB.append( input() )
    # end for

    for word in groupB:
        # if the word exists in groupA
        if( word in groupA ):
            # show positions
            print( *groupA[word] )
        else:
            # else, print -1
            print( "-1" )
