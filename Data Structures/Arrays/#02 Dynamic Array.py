if __name__ == '__main__':

    # numSeq -> numero of sequences
    # numQue -> numero of queries
    numSeq, numQue = map( int, input().split() )

    # sequences
    seqList = list()

    # make N sequences
    for _ in range (numSeq) :
        seqList.append( list() )
    # end for

    lastAnswer = 0

    # for each query
    for _ in range (numQue):

        # type of query: 1 or 2
        # x and y from problem
        typeQuery, x, y = map( int, input().split() )

        if( typeQuery == 1 ):
            # find the sequence seq, at index ( (x^lastAnswer) % numSeq ) in seqList
            seq = ( ( x ^ lastAnswer ) % numSeq )
            # append integer y to sequence seq
            seqList[ seq ].append( y )
        else:
            # find the sequence, seq, at index( (x^lastAnswer) % numSeq ) in seqList
            seq = ( ( x ^ lastAnswer ) % numSeq )
            # find the value of element y % size in seq( where size is the size of seq )
            # and assign it to lastAnswer
            value = y % len( seqList[seq] )
            lastAnswer = seqList[ seq ][ value ]
            # print the new value of lastAnswer on a new line
            print( lastAnswer )
    # end for
