from collections import OrderedDict

def mean( elements, numElements ):
    return ( sum(elements) / numElements )
# end mean

def median( elements, numElements ):
    if( numElements % 2 == 0 ):
        return ( ( elements[numElements//2] + elements[numElements//2-1] ) / 2 )
    else:
        return ( elements[numElements//2] )
# end median
      
def mode ( elements, numElements ):
    count = OrderedDict()
      
    for e in elements:
        if( e in count.keys() ):
            count[e] += 1
        else:
            count[e] = 1
    # end for

    maxValue = max( count.values() )
      
    count = [(x,count[x]) for x in count]
    count.sort( key=lambda x: x[1] )
    
    for x in count:
        if( x[1] == maxValue ):
            return x[0]
# end mode  

if __name__ == '__main__':
    # number of elements
    numElements = int( input() )

    # read all elements as integer and store in a list
    elements = list( map(int, input().split()) )
    # sort the list
    elements.sort()

    # mean
    print( "%.1f"%( mean(elements,numElements) ) )
    # median
    print( "%.1f"%( median(elements,numElements) ) )
    # mode
    print( mode(elements,numElements) )
