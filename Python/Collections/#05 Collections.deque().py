from collections import deque

if __name__ == '__main__':

    d = deque()

    for _ in range( int(input()) ):
        command = input().split()

        if( command[0] == "append" ):
            d.append( int( command[1] ) )
        elif( command[0] == "pop"):
            d.pop()
        elif( command[0] == "popleft"):
            d.popleft()
        elif( command[0] == "appendleft"):
            d.appendleft( int( command[1] ) )
    # end for

    print( *d )
