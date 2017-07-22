#!/bin/python3

import sys

def isLuckyNumber( number ):
  number = str(number)
  return( sum( map(int,number[0:3])) == sum( map(int,number[3:6])) )

def onceInATram(x):
    x += 1
    while( isLuckyNumber(x) == False ):
      x +=1
    return x

if __name__ == "__main__":
    x = int(input().strip())
    result = onceInATram(x)
    print(result)
