if __name__ == '__main__':

    # first number
    a = int( input() )
    # second number
    b = int( input() )

    division = divmod( a,b )

    print( *[ division[0], division[1], division ], sep = '\n' )
