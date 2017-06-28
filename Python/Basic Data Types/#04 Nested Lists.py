if __name__ == '__main__':

    students = list()
    grades = set()

    for _ in range(int(input())):
        name = input()
        score = float(input())

        grades.add( score )
        students.append( (name,score) )
    # end for

    grades = list( grades )
    grades.sort()

    # sort by name
    students.sort( key = lambda x: x[0] )

    print( *[ x[0] for x in students if( x[1] == grades[1] ) ], sep = '\n')
