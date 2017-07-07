if __name__ == '__main__':

    # the N from input is 0 < N < 100, then i can get all palindromic numbers
    palindromic = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99}

    # i don't need this number
    input()

    # get numbers from input
    numbers = list( map(int, input().split() ) )

    # test if any number in input is palindromic, and
    # test if all numbers are positive numbers
    print( any( [ n in palindromic for n in numbers ] ) and all( [ (n>0) for n in numbers ] ) )

'''
My 3 lines solution:

palindromic = {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 22, 33, 44, 55, 66, 77, 88, 99}
input() ; numbers = list( map(int, input().split() ) )
print( any( [ x in palindromic for x in numbers ] ) and all( [ x > 0 for x in numbers ] ) )
'''
