'''
This chalenge is only for Python2
'''

if __name__ == '__main__':

    # x tem que ter o mesmo nome da variavel do input
    x, k = map( int, raw_input().split() )

    # opera com o x lido acima
    func = input()

    print func == k

'''
Python3 equivalent

Author: qingchen

x,k=map(int, input().split())
print (k==eval(input()))
'''
