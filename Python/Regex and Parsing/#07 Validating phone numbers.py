import re

# [789] = (7|8|9)
pattern = "[789]\d{9}$"

# for each test case
for _ in range( int(input()) ):
    number = input()

    if( re.match(pattern, number) ):
        print( "YES" )
    else:
        print( "NO" )
