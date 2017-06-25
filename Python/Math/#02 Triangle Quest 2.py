for i in range(1,int(input())+1): #More than 2 lines will result in 0 score. Do not leave a blank line also
    print( pow( (pow(10,i)-1)//9 , 2 ) )

'''
In the internet i founded this propriety:
1x1 = 1
11x11 = 121
111x111 = 12321
1111x1111 = 1234321

multiply 111x111:

    111
   x111
   ---
    111
   111
  111
   ---
  12321

*******************************************************

I got this:
111 = 1x10^2 + 1x10^1 + 1x10^0

then

111 = { sum 10^i, i=1 to n-1 } ; wolfram alpha notation

then the formula is
( 10^(n-1) )/9 * ( 10^(n-1) )/9
'''
