if __name__ == '__main__':
    n = int( input() )

    # Author: SemperPeritus
    # For short operator '*' allows you to pass an array
    # sep is the separator
    print( *range(1,n+1), sep='' )
