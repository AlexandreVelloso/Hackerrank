import re

# for each test case
for _ in range( int( input() ) ):
    name, email = input().split()
    # remove the first, and the last character, (<,>)
    email = email[ 1 : -1 ]

    pattern = "[a-z](\w|\.|-)*@[a-z]+\.[a-z]{1,3}$"

    if( re.match( pattern, email) ):
        print( name, "<"+email+">" )
    # end if
# end for
