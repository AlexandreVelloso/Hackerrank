if __name__ == '__main__':

    # i don't need this line
    input()

    # english newspapper
    english = set( map( int, input().split() ) )

    # i don't need this line
    input()

    # french newspapper
    french = set( map( int, input().split() ) )

    print( len( english.union( french ) ) )
