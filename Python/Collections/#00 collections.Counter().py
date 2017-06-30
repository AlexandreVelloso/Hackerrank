if __name__ == '__main__':

    # i don't need the fisrt line
    input()

    # list of shoe sizes
    listSizes = list( map( int, input().split() ) )

    # money earned by Raghu
    sum = 0

    # for each client
    for _ in range( int(input()) ):

        # each client -> [ shoe size, price ]
        size, price = map( int, input().split() )

        # if this size is avaliable
        if( size in listSizes ):

            # sell the shoe
            sum += price
            # remove from stock
            listSizes.remove( size )
    # end for

    print( sum )

'''
I didn't use collections, this is the problem solved using collections

Author: DOSHI

from collections import Counter
X = input()
S = Counter(map(int,raw_input().split()))
N = input()
earnings = 0
for customer in range(N):
    size, x_i = map(int,raw_input().split())
    if size in S and S[size] > 0:
        S[size] -= 1
        earnings += x_i

print earnings
'''
