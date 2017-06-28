cube = lambda x: x**3

def fibonacci(n):

    # primeiros 3 fibonaccis conhecidos
    fib = [0,1,1]

    # calcula o fibonacci dos numeros ate n
    for x in range(3,n):
        fib.append( fib[x-1] + fib[x-2] )

    # retorna o fibonacci dos numeros ate n
    return [ fib[x] for x in range(n) ]

if __name__ == '__main__':
    n = int(input())
    print(list(map(cube, fibonacci(n))))
