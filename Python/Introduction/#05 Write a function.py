def is_leap(year):
    leap = False

    if( year%100 != 0 and year%4 == 0 ):
        leap = True
    elif( year%100 == 0 and year%400 == 0 ):
        leap = True

    return leap

#end is_leap

if __name__ == '__main__':

    year = int(input())
    print(is_leap(year))

# end main
