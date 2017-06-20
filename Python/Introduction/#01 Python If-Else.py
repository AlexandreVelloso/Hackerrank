if __name__ == '__main__':

    # read int from stdin
    n = int( input() )

    # if id odd
    if( n % 2 == 1 ):
        print( "Weird" )
    elif( n >= 2 and n <= 5 ):
        print( "Not Weird" )

    elif( n in range(6,21) ):
        print( "Weird" )

    elif( n > 20 ):
        print( "Not Weird" )
#end main


# author: Diapolo10
'''
n = int(input().strip())

# We define a dictionary
check = {True: "Not Weird", False: "Weird"}

print( check[ n%2==0 and ( n in range(2,6) or n > 20) ] )
'''
