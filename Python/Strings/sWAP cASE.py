def swap_case(s):

    swaped = ""

    for c in s:
        # if uppercase
        if( c >= 'A' and c <= 'Z' ):
            swaped += c.lower()

        # if lowercase
        elif( c>= 'a' and c <= 'z' ):
            swaped += c.upper()
            
        else:
            swaped += c
    # end for

    return swaped
# end swap_case

if __name__ == '__main__':
    s = input()
    result = swap_case(s)
    print(result)
