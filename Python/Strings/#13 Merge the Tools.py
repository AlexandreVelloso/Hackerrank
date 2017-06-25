def merge_the_tools(string, k):

    for i in range( 0, len(string), k ):

      u = ""

      for x in range( i, i+k ):
        if( not (string[x] in u ) ):
            u += string[x]

      print( u )

if __name__ == '__main__':
    string, k = input(), int(input())
    merge_the_tools(string, k)
