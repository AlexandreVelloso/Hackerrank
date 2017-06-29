import re

if __name__ == '__main__':

    for x in range( int(input() ) ):
        str = input()

        print( bool( re.match("^[+-]?[0-9]*\.[0-9]+$", str) ) )

'''
explanation:
^       begin of the string
[+-]?   opcional x,-,lambda
[0-9]*  0 or more digits
\.      character .
[0-9]+  1 or more digits
$       end of the string
'''
