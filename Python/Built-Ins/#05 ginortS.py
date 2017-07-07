'''
My solution is based on stefan_pochmann solution

print(*sorted(input(), key=lambda c: (c.isdigit() - c.islower(), c in '02468', c)), sep='')
'''
if __name__ == '__main__':
    S = input()

    # ordena a string pela ordem de precedencia,
    # quanto mais a direita a condicao esta, maior precedencia ela tem
    S = sorted(S, key = lambda c: ( c in "02468", c.isdigit(), c.isupper(), c.islower(), c ) )

    print( *S, sep = '' )

'''
Solution by editorial:

Solution 1: using key

from __future__ import print_function

def func(x):
    if x.isalpha():
        if x.isupper():
            return (ord(x)-ord('A'))
        else:
            return (ord(x)-ord('a'))-30
    else:
        if int(x) % 2 == 0:
            return 60+int(x)
        else:
            return 30+int(x)

s = raw_input()
map(lambda x: print(x,end=''),(sorted(s,key = func)))

*********************************************************

Solution2: Without using key

from __future__ import print_function

upper = []
lower = []
even = []
odd = []

def separator(a):

    if a.isalpha():
        if a.isupper():
            upper.append(a)
        else:
            lower.append(a)
    else:
        if int(a)%2 == 0:
            even.append(a)
        else:
            odd.append(a)
    return

map(separator,raw_input())

upper.sort()
lower.sort()
even.sort()
odd.sort()

t = lower+upper+odd+even
map(lambda x: print(x,end=''),t)
'''
