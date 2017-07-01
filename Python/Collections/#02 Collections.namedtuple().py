from collections import namedtuple

# number of students
N = int( input() )

Student = namedtuple('Student', input() )

# sum of the students' marks
sumMarks = 0

for _ in range( N ):
    # read a student
    std = Student( *input().split() )
    # sum his MARKS
    sumMarks += int( std.MARKS )

# print the avarage
print ( "%.2f" % (sumMarks / N) )
