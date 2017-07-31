import sys

numCities, numStations = map( int, input().split() )
stations = list( map( int, input().split() ) )

# sort the stations first
stations.sort()

# biggest distance founded
maxDistance = 0

# special cases
if( len(stations) >= 2 ):
    i = stations[0]
    j = stations[1]
    k = 1
elif( len(stations) == 1 ):
    i = stations[0]
    j = i
    k = 0
else:
    # 0 stations
    i,j,k = 0,0,0
   
# for every city, get the minimum distance for a station
for x in range( numCities ):
    
    if( x <= i ):
        
        distance = abs( i-x )
        #
        maxDistance = max( maxDistance, distance )
    elif( x <= j ):
        
        distance = min( abs( i-x ), abs( j-x ) )
        #
        maxDistance = max( maxDistance, distance )
    else:
        if( k+1 < numStations ):
            i = j
            k += 1
            j = stations[k]
        # end if
            
        distance = min( abs( i-x ), abs( j-x ) )
        #
        maxDistance = max( maxDistance, distance )
    # end if
    
# end for
print( maxDistance )