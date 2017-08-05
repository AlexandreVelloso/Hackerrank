import numpy

def arrays(arr):

    # function to flip a array
    flip = numpy.fliplr([arr])[0]

    # transform the int array to a float array, then returns
    return( numpy.array( flip, float ) )
# end function arrays

arr = input().strip().split(' ')
result = arrays(arr)
print(result)
