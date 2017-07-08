import re

if __name__ == '__main__':

    s = input()

    vowels = "aeiou"
    consoants = "qwrtypsdfghjklzxcvbnm"

    # (?<=[consoants]) verifica se a string contem o prefixo, sem consumir a string
    # ([vowels]{2,})   verifica se a string contem 2 ou mais vogais
    # (?=[consoants])  verifica se a string contem o sufixo, consumindo a string
    p = "(?<=[%s])([%s]){2,}(?=[%s])" % ( consoants, vowels, consoants )

    # compila o padrao, a parte re.I e' ignore case
    pattern = re.compile( p, re.I )

    # acha os padroes que casam com o pattern
    result = pattern.finditer( s )

    # essa gambiarra eu fiz porque foi a solucao que eu encontrei pra quando nao achar o padrao
    if( len( pattern.findall(s) ) ):
        # mostra todos os resultados se encontrados
        print( *map( lambda x: x.group() if result else -1, result ), sep = '\n' )
    else:
        # mostra -1 se nao encontrou nenhum padrao
        print(-1)
