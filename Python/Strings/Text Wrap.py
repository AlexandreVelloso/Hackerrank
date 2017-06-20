import textwrap

def wrap(string, max_width):
    return textwrap.fill( string, max_width )

if __name__ == '__main__':
    string, max_width = input(), int(input())
    result = wrap(string, max_width)
    print(result)

'''
Author: shanker4999

without using textwrap
s=raw_input().strip()
w=int(raw_input())
for i in range(0,len(s)+1,w):
    print(s[i:w+i])
'''
