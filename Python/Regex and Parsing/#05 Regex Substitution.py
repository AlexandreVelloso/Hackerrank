import re

# function to replace:
# && -> and
# || -> or
def funcReplace(match):
    simbol = match.group(0)

    if( simbol == '&&' ):
        return "and"
    elif( simbol == '||' ):
        return "or"
# end funcReplace

if __name__ == '__main__':

    strPattern = "(?<= )({}|{})(?= )".format( "&&", "\|\|" )
    pattern = re.compile( strPattern )

    # for each line from input
    for _ in range( int(input()) ):
        print( pattern.sub( funcReplace, input() ) )

# end main
