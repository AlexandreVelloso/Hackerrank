from collections import Counter

if __name__ == '__main__':

    '''
    # i don't use K
    K = int( input() )

    # read the array from input
    array = map(int, input().split() )

    # count the number of occurrence of each number
    count = Counter( array )

    # print the min number
    print( min( count, key = count.get ) )
    '''

'''

    Solution from the editorial, i don't get it

    K = int(input())
    set_S = set()
    sumlist_S = 0
    for i in input().split():
        I = int(i)
        set_S.add(I)
        sumlist_S += I

    print((sum(set_S)*K - sumlist_S)//(K-1))
'''
