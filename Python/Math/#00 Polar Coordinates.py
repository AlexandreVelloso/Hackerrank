import cmath

if __name__ == '__main__':
    number = complex( input() )
    print( abs( number ) )
    print( cmath.phase(number) )
