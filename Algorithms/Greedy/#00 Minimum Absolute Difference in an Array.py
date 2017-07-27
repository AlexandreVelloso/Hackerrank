if __name__ == "__main__":
    # number of elements in the array
    numElements = int( input() )
    # array of elements
    array = list( map(int, input().split()) )
    
    # sort the array
    array.sort()
    
    # initialize the minimum diference
    minimumDifference = abs( array[0] - array[1] )
    
    # find the minimum diference in the array
    for x in range( len(array) - 1 ):
        minimumDifference = min( minimumDifference, abs( array[x] - array[x+1] ) )
    # end for
    
    print( minimumDifference )