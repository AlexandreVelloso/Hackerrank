from math import sqrt

def mean( array ):
    return ( sum(array) / len(array) )
# end mean

def standartDeviation ( array ):

    meanArray = mean( array )

    numerator = sum( [(x - meanArray)**2 for x in array ] )
    return sqrt( numerator / len( array) )

# end methos standartDeviation

if __name__ == '__main__':
    # number of elements
    numElements = int( input() )
    # array of elements
    array = list( map(int, input().split() ) )

    print ( "%.1f" % ( standartDeviation( array ) ) )
