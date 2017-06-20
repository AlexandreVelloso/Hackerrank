if __name__ == '__main__':

    # number of cases
    n = int( input() )
    marks = {}

    for _ in range( n ):
        student = input().split()

        # student average
        average = ( float(student[1]) + float(student[2]) + float(student[3]) ) / 3.0
        # map name to average
        marks[ student[0] ] = average

    # print the average
    print( "{0:.2f}".format( marks[ input() ] ) )
