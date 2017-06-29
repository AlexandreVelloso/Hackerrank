from itertools import combinations

# number of letters
N = int( input() )
letters = list( input().split() )
# number of letters to pe picked
K = int( input() )

# make all combinations
comb = list( combinations( letters, K ) )
# count the number of letters a in the combination
number_a = ['a' in x for x in comb ].count(True)

print( number_a / len(comb) )
