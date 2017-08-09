import re

# this pattern i get from discutions
# I don't know why this works because of [\s:]
# \s matches any whitespace character (equal to [\r\n\t\f\v ])
# : matches the character : literally (case sensitive)
pattern = "[\s:](#[a-fA-F0-9]{6}|#[a-fA-F0-9]{3})"
pattern = re.compile( pattern, re.I )

# for each line
for _ in range( int(input()) ):
    line = input()

    for p in pattern.findall( line ):
        print( p )

# end for
