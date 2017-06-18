if __name__ == '__main__':
    n = int( input() )

    lista = list()

    # for each comand
    for _ in range(n):

        command = input().split()

        if( command[0] == "insert" ):

            i = int( command[1] )
            e = int( command[2] )
            # insert integer 'e' in position 'i'
            lista.insert( i, e )

        elif( command[0] == "print"):

            print( lista )

        elif( command[0] == "remove" ):

            e = int( command[1] )
            lista.remove( e )

        elif( command[0] == "append" ):

            e = int( command[1] )
            lista.append( e )

        elif( command[0] == "sort" ):

            lista.sort()

        elif ( command[0] == "pop" ):

            lista.pop()

        elif( command[0] == "reverse" ):

            lista.reverse()

#end main

'''
Author: dennysregalado

n = input()
l = []
for _ in range(n):
    s = raw_input().split()
    cmd = s[0]
    args = s[1:]
    if cmd !="print":
        cmd += "("+ ",".join(args) +")"
        eval("l."+cmd)
    else:
        print l
'''
