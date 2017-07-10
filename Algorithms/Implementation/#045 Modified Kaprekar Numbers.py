def isKaprekarNumber( num ):

    if( num < 1 ):
        return False
    elif( num == 1 ):
        return True
    else:
        power = num**2

        # convert the power to string
        powerAsString = str(power)
        # number of digits of power
        numDigits = len(powerAsString)

        # left part of the number
        left = powerAsString[ 0 : numDigits//2]
        # right part of the number
        right = powerAsString[ numDigits//2 : numDigits]

        if( ( len(left) > 0 ) and ( len(right) > 0) ):
            # divide the number in to parts, and sum then
            sumParts = int( left ) + int( right )
        else:
            sumParts = 0

        return ( num == sumParts )

# end function isKaprekarNumber

if __name__ == '__main__':

    # get all kaprekar numbers in the inclusive range (p,q)
    numbers = [ x for x in range( int(input()), int(input()) + 1 ) if( isKaprekarNumber(x) ) ]

    if( len(numbers) > 0 ):
        print( *numbers, sep = ' ')
    else:
        print( "INVALID RANGE" )
