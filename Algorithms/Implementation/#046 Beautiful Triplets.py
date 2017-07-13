if __name__ == '__main__':

    # size -> size of the sequence
    # d    -> constant from the problem
    size, d = map (int, input().split())
    sequence = list( map( int, input().split() ) )

    count = 0

    for i in range( size ):
        # sequence[j] = (sequence[i] + d) in sequence
        # sequence[k] = (2*d + sequence[i]) in sequence
        if( ( (sequence[i] + d) in sequence ) and ( (2*d + sequence[i]) in sequence) ):
            count +=1
    # end for

    print( count )

    '''

    A beautiful triplet is

    a[j] - a[i] = a[k] - a[j] = d

    Then, the element j is:

    a[j] - a[i] = d
    a[j] = d + a[i]

    And the element k is:

    a[k] - a[j] = d
    a[k] = d + a[j]
    *** a[j] = d + a[i] ***
    a[k] = d + d + a[i]
    a[k] = 2*d + a[i]
    
    '''
