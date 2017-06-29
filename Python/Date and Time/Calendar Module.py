import calendar

month, day, year = map( int, input().split() )

dayOfWeek = [ "MONDAY","TUESDAY","WEDNESDAY","THURSDAY","FRIDAY","SATURDAY","SUNDAY" ]

print( dayOfWeek [ calendar.weekday( year, month, day) ] )
