for _ in range( int(input()) ):

    try:
        num1, num2 = map( int, input().split() )
        print( num1 // num2 )

    except ValueError as e:
        print ( "Error Code:",e )
    except ZeroDivisionError as e2:
        print ( "Error Code: integer division or modulo by zero" )
