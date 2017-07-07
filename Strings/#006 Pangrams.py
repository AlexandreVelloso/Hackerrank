from collections import Counter

if __name__ == '__main__':

    # read input, remove the spaces, and convert the string to lowercase
    phase = input().replace(' ','').lower()

    # count the number of occurrences for each character
    count = Counter( phase )

    if( len(count) == 26 ):
        print( "pangram" )
    else:
        print( "not pangram" )
