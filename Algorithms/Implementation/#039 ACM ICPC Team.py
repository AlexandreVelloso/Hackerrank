from collections import Counter

# count how many topics a team know
def countTopics( people1, people2 ):
    count = 0

    for i in range( len(people1) ):
        if( people1[i] == '1' or people2[i] == '1' ):
            count += 1

    # end for

    return count
# end method

if __name__ == '__main__':

    # N and M from problem
    numPeople, numTopics = map( int, input().split() )

    # peoples from input data
    peoples = []
    # number of topics for each team
    numTopics = []

    # read all peoples from input
    for _ in range( numPeople ):
        peoples.append( input() )

    # for each pair (i,j) where i != j
    for i in range( numPeople ):
        for j in range( i+1, numPeople ):
            numTopics.append( countTopics( peoples[i], peoples[j] ) )

    # find the maximun number of topics that a team knows
    maximum = max( numTopics )
    print( maximum )

    # count the number of occurrences from each number of topics
    counter = Counter( numTopics )
    # print the number of teams who knows the maximum number of topics
    print( counter[maximum] )
