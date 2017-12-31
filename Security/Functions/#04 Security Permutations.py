array = []

def f( x ):
    return array[ x-1 ][0]

if __name__ == '__main__':

  # number of elements
  n = int( input() )

  x = 1
  # for each element, make a map ( x -> e )
  for e in map( int, input().split() ):
      array.append( (x,e) )
      x +=1
  # end for

  # sort array by second element in tuple
  array.sort( key=lambda x: x[1] )

  print( array )

  print( f(1) )
  print( f(2) )
  print( f(3) )

  # print the elements in order
  #print( *[ f[ f[ x-1 ][1] -1][1] for x in range( n ) ], sep = '\n' )
