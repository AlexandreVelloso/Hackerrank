import sys
if __name__ == '__main__':
    
    numCap, numMax = map( int, input().split() )
    numProbCap = list( map(int, input().split() ) )
    
    pageNumber = 1
    
    numberSpecial = 0
    
    for nump in numProbCap:
        
        x = 1
        # paginas completas
        for _ in range( nump//numMax ):
            
            # gera o numero dos exercicios dessa pagina
            lista = [ x for x in range( x, numMax+x )]
            x += numMax
            
            if( pageNumber in lista ):
                numberSpecial += 1
                
            pageNumber += 1
        # end for
        
        # paginas incompletas
        if( nump%numMax > 0 ):
            
            # gera o numero dos exercicios dessa pagina
            lista = [ x for x in range( x, (nump%numMax)+x )]
            
            if( pageNumber in lista ):
                numberSpecial += 1
            
            pageNumber += 1
        # end if

    # end for
    
    print( numberSpecial )
# end main