if __name__ == '__main__':

    hour = int( input() )
    minutes = int( input() )

    numerals = { 1:"one", 2:"two", 3:"three", 4:"four", 5:"five",
                 6:"six", 7:"seven", 8:"eight", 9:"nine", 10:"ten",
                 11:"eleven", 12:"twelve", 13:"thirteen", 14:"fourteen", 15:"fifteen",
                 16:"sixteen", 17:"seventeen", 18:"eighteen",19:"nineteen", 20:"twenty" }

    if( minutes == 0 ):
        print( numerals[hour]+" o' clock" )
    elif( minutes == 15 ):
        print( "quarter past "+numerals[hour] )
    elif( minutes == 30 ):
        print( "half past "+numerals[hour] )
    elif( minutes == 45 ):
        print( "quarter to "+ (numerals[ (hour+1) % 12 ] ) )
    elif( minutes > 30 ):
        minutes = 60 - minutes

        converted = ""

        decimal = minutes // 10
        unit = minutes % 10

        if( decimal == 1 ):
            converted += numerals[ minutes ]
        elif( decimal == 2 ):
            converted += numerals[ 20 ]+" "
            converted += numerals[ unit ]
        else:
            converted += numerals[ unit ]

        print( converted +" minutes to "+numerals[ (hour+1) % 12 ] )
    else:

        converted = ""

        decimal = minutes // 10
        unit = minutes % 10

        if( decimal == 2 ):
            converted += numerals[ 20 ]+" "
        converted += numerals[ unit ]

        print( converted +" minutes past "+ numerals[ hour ] )
