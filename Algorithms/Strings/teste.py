def f( numStr, inicio, fim, ant ):

    tam = len( numStr )

    if( fim < tam ):

        num = numStr[inicio:fim+1]
        print( "n:%s a:%s" % (num,ant) )

        f( numStr, inicio , fim + 1, num )

    elif( inicio <= tam ):

        num = numStr[inicio:fim+1]
        print( "n:%s a:%s" % (num,ant) )

        f( numStr, inicio + 1, inicio+1, num )
# end method f

# main
f( "1234",0,0, "" )
# end main
