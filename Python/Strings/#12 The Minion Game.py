if __name__ == '__main__':
    word = input()

    vowels = 'AEIOU'

    kevinScore = 0
    stuartScore = 0

    for index in range( len(word) ):
        if( word[index] in vowels ):
            kevinScore += len(word) - index
        else:
            stuartScore += len(word) - index

    if( kevinScore > stuartScore ):
        print( "Kevin", kevinScore )
    else:
        print( "Stuart", stuartScore )

'''
Explanation by:

It's an implicit (and smart) count of occurrence.

Es.

BANANA (for kevin):

For i=1 we are sure that there is an occurrence for all this substrings that start with a vowel (s[i] = 'A'):

- A
- AN
- ANA
- ANAN
- ANANA
So we add len(s)-i = 6-1 = 5 to kevin score.

For i=3:

- A
- AN
- ANA
(+3)

For i=5:

- A
(+1)
'''
