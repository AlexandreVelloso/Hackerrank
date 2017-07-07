if __name__ == '__main__':

    numStudents, numSubjects = map( int, input().split() )

    marks = list()

    # read the input
    for _ in range( numSubjects ):
        marks.append( map (float, input().split() ) )
    # end for

    for std_marks in zip(*marks):
        print( sum(std_marks) / numSubjects )
    # end for
# end main

'''
Based on the Editorial
'''
