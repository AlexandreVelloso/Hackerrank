from datetime import datetime as dt

# Author: FJSevilla

fmt = '%a %d %b %Y %H:%M:%S %z'

for _ in range( int(input())):
  data1 = dt.strptime( input(), fmt )
  data2 = dt.strptime( input(), fmt )

  print( int( abs( (data2 - data1).total_seconds() ) ) )
