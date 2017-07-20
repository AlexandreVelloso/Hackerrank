# size    -> size of the freeway
# numTest -> number of test cases
size, numTest = map( int, input().split() )

# widths of the services lanes
width = list( map(int, input().split()) )

# for each test case
for _ in range( numTest ):
  # i-> entry
  # j-> exit
  i, j = map( int, input().split() )
  
  # find the minimum length of the service lane  
  print( min( width[i:j+1] ) )
