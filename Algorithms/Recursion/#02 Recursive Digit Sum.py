def superDigit( n ):
    
    if( len(n) > 1 ):
        
        # sum of all digits
        sumDigits = sum( [ int(digit) for digit in n ] )
        
        # transform the sum to string, and get this superDigit
        return( superDigit( str(sumDigits) ) )
    else:
        return (n)
# end method superDigit

if __name__ == "__main__"
    # read n and k
    lineSplited = input().split()

    # first, read n is a string
    n = lineSplited[0]
    # read k as integer
    k = int( lineSplited[1] )

    # sum the digits of n
    sumDigits = sum( [ int(digit) for digit in n ] )

    # multiply the sum by k, and begin the recursion
    print( superDigit( str(sumDigits * k) ) )
