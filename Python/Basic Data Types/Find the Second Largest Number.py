if __name__ == '__main__':
    n = int(input())
    arr = list( map(int, input().split() ) )

    first = max( arr[0], arr[1] )
    second = min( arr[0], arr[1] )

    for x in range(2,n):

        if( arr[x] > first ):
            second = first
            first = arr[x]
        elif( arr[x] != first and arr[x] > second ):
            second = arr[x]
        elif( first == second ):
            second = min( second, arr[x] )
    # end for

    # print the result
    print( second )

# end main
