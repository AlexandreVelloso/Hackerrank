import re

for _ in range( int(input() ) ):

    try:
        print( bool( re.compile( input() ) ) )
    except (Exception):
        print("False")
