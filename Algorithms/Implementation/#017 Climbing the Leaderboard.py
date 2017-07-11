from collections import defaultdict

# map ( score -> rank )
ranks = defaultdict( list )

def binarySearch( element, array ):

    left = 0
    right = len(array) - 1

    while( left <= right ):
        middle = (left+right) // 2

        if( array[middle] == element ):
            return ( ranks[ element ] )
        elif( array[middle] > element ):
            left = middle + 1
        else:
            right = middle - 1
    #end while

    if( array[right] > element ):
        return ( ranks[ array[right] ]+1 )
    else:
        return 1
# end functionn binarySearch

if __name__ == '__main__':

    # number of players already in the leaderboard
    numPlayers = int( input() )
    # playes' scores
    playersScores = list( map( int, input().split() ) )
    # number of levels that Alice beats
    numLevels = int( input() )
    # Alice's scores
    aliceScores = list( map( int, input().split() ) )

    # ****** fill the ranks ******

    # this is a map between score and rank

    lastRank = 0
    prevous = -1
    for score in playersScores:
        if( score != prevous ):
            lastRank += 1
            prevous = score

        ranks[ score ] = lastRank

    # ****** end fill the ranks ******

    for aliceS in aliceScores:
        print( binarySearch( aliceS, playersScores) )
