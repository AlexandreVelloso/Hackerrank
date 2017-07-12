if __name__ == '__main__':

  # number of elements
  n = int( input() )

  array = []

  x = 1
  # for each element, make a map ( x -> e )
  for e in map( int, input().split() ):
      array.append( (x,e) )
      x +=1
  # end for

  # sort array by second element in tuple
  array.sort( key=lambda x: x[1] )

  # print the elements in order
  print( *[ x[0] for x in array ], sep = '\n' )
