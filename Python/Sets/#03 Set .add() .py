if __name__ == '__main__':
    # number of stumps
    n = int( input() )
    stumps = set()

    for _ in range(n):
        stumps.add( input().strip().upper() )

    print( len(stumps) )
