if __name__ == '__main__':
    
    # size of the array
    n = int( input() )
    # array of elements
    array = list( map(int, input().split() ) )
    # array of weights
    weights = list( map(int, input().split() ) )
    
    
    numerator = sum( [ (array[x] * weights[x]) for x in range( n ) ] )
    denominator = sum( weights )
    
    print( "%.1f"%(numerator/denominator) )
