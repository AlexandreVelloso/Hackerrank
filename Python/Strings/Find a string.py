def count_substring(string, sub_string):

    count = 0

    # the len(string) is <= 200, then i don't need to worry
    # about efficience here
    for pos in range( 0, len( string ) - len( sub_string ) + 1 ):

        if( string[ pos : pos+len(sub_string) ] == sub_string ):
            count += 1

    return count
# end count_substring

if __name__ == '__main__':
    string = input().strip()
    sub_string = input().strip()

    count = count_substring(string, sub_string)
    print(count)
