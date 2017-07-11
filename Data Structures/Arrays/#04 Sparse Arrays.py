from collections import Counter

if __name__ == '__main__':

    # numer of strings
    numStr = int( input() )

    # list of strings
    strList = list()

    # read the string fron input
    for _ in range( numStr ):
        strList.append( input().strip() )
    # enf for

    # count the number of occurrences from each string
    counter = Counter( strList )

    # number of queries
    numQue = int( input() )

    # for each query
    for _ in range( numQue ):
        # print the occurrence from each string from input()
        print( counter[input().strip()] )
