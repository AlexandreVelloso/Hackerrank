import math

if __name__ == '__main__':

    AB = int( input() )
    BC = int( input() )

    # hypotenuse
    AC = math.sqrt( AB*AB + BC*BC )

    # midpoint
    MC = AC/2.0

    angle = math.degrees( math.acos( (BC/2.0)/MC ) )

    print( str( round(angle) ) + "°")

'''

Found the hypotenuse AC
Fount the midpoint MC

To calculate the angle i used the acos( midpoint(BC)/MC )

'''
