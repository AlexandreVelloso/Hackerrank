#!/bin/python3

import sys

def lonely_integer(a):
  value = 0
  
  for val in a:
    value ^= val
    
  return value
    
# end lonely_integer

n = int(input().strip())
a = [int(a_temp) for a_temp in input().strip().split(' ')]
print(lonely_integer(a))

'''
*bonillasc :
  I don't understand the solution. Take an array of int's [a a b b c d d]

  Now the solution posted will run like this 0 ^ a = a, a ^ a = 0 (two pairs of a's cancel out, same with b)

    but then we arrive at c; where value is 0. 0^c = c, BUT THEN c ^ d = c^d, then (c^d) ^ d = (c^d) ^ d = value at the end of this array.

  So function returns (c^d) ^ d What am I misunderstanding?
  
*superrap4761 :
  xor operation is both commutative and associative. This means that, 1. a ^ b = b ^ a (Commutative property) 2. a ^ (b ^ c) = (a ^ b) ^ c (Associative property)

  So according to your array, the final value is, value = (c ^ d) ^ d which, by Associative property(2), evaluates to, value = c ^ (d ^ d). But (d ^ d) = 0. Hence, value = c ^ (0) = c , the lonely integer
'''
