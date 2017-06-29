import re

if __name__ == '__main__':

    for x in range( int(input() ) ):
        number = input()

        print( bool( re.match("^[+-]?[0-9]*\.[0-9]+$", number) ) )

'''
explanation:
^       begin of the string
[+-]?   opcional x,-,lambda
[0-9]*  0 or more digits
\.      character .
[0-9]+  1 or more digits
$       end of the string
'''
